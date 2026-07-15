# ====================================================================================
# Setup Project

PROJECT_NAME ?= provider-gridscale
PROJECT_REPO ?= github.com/PlatformRelay/$(PROJECT_NAME)

export TERRAFORM_VERSION ?= 1.5.7

# Do not allow a version of terraform greater than 1.5.x, due to versions 1.6+ being
# licensed under BSL, which is not permitted.
TERRAFORM_VERSION_VALID := $(shell [ "$(TERRAFORM_VERSION)" = "`printf "$(TERRAFORM_VERSION)\n1.6" | sort -V | head -n1`" ] && echo 1 || echo 0)

export TERRAFORM_PROVIDER_SOURCE ?= gridscale/gridscale
export TERRAFORM_PROVIDER_REPO ?= https://github.com/gridscale/terraform-provider-gridscale
export TERRAFORM_PROVIDER_VERSION ?= 2.3.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME ?= terraform-provider-gridscale
export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX ?= $(TERRAFORM_PROVIDER_REPO)/releases/download/v$(TERRAFORM_PROVIDER_VERSION)
export TERRAFORM_NATIVE_PROVIDER_BINARY ?= terraform-provider-gridscale_v2.3.0
export TERRAFORM_DOCS_PATH ?= website/docs/r


PLATFORMS ?= linux_amd64 linux_arm64

# -include will silently skip missing files, which allows us
# to load those files with a target in the Makefile. If only
# "include" was used, the make command would fail and refuse
# to run a target until the include commands succeeded.
-include build/makelib/common.mk

# ====================================================================================
# Setup Output

-include build/makelib/output.mk

# ====================================================================================
# Setup Go

# Set a sane default so that the nprocs calculation below is less noisy on the initial
# loading of this file
NPROCS ?= 1

# each of our test suites starts a kube-apiserver and running many test suites in
# parallel can lead to high CPU utilization. by default we reduce the parallelism
# to half the number of CPU cores.
GO_TEST_PARALLEL := $(shell echo $$(( $(NPROCS) / 2 )))

GO_REQUIRED_VERSION ?= 1.24
GOLANGCILINT_VERSION ?= 2.12.1
GO_STATIC_PACKAGES = $(GO_PROJECT)/cmd/provider $(GO_PROJECT)/cmd/generator
GO_LDFLAGS += -X $(GO_PROJECT)/internal/version.Version=$(VERSION)
# config holds the hand-authored provider configuration (external-name strategy
# and kind overrides) and its unit tests. It must be in GO_SUBDIRS so that
# go.test.unit (which tests $(GO_PACKAGES) derived from GO_SUBDIRS) actually
# compiles and runs config/*_test.go and reports its coverage.
GO_SUBDIRS += cmd internal apis config
-include build/makelib/golang.mk

# ====================================================================================
# Setup Kubernetes tools

KIND_VERSION = v0.31.0
UPTEST_VERSION = v2.2.0
CRDDIFF_VERSION = v0.12.1
CROSSPLANE_CLI_VERSION = v2.2.1
# for e2e testing
CROSSPLANE_VERSION = 2.2.1
-include build/makelib/k8s_tools.mk

# ====================================================================================
# Setup Images

REGISTRY_ORGS ?= ghcr.io/platformrelay
IMAGES = $(PROJECT_NAME)
-include build/makelib/imagelight.mk

# ====================================================================================
# Setup XPKG

XPKG_REG_ORGS ?= ghcr.io/platformrelay
# NOTE(hasheddan): skip promoting on xpkg.crossplane.io as channel tags are
# inferred.
XPKG_REG_ORGS_NO_PROMOTE ?= ghcr.io/platformrelay
XPKGS = $(PROJECT_NAME)
-include build/makelib/xpkg.mk

# ====================================================================================
# Fallthrough

# run `make help` to see the targets and options

# We want submodules to be set up the first time `make` is run.
# We manage the build/ folder and its Makefiles as a submodule.
# The first time `make` is run, the includes of build/*.mk files will
# all fail, and this target will be run. The next time, the default as defined
# by the includes will be run instead.
fallthrough: submodules
	@echo Initial setup complete. Running make again . . .
	@make

# NOTE(hasheddan): we force image building to happen prior to xpkg build so that
# we ensure image is present in daemon.
xpkg.build.provider-gridscale: do.build.images

# NOTE(hasheddan): we ensure up is installed prior to running platform-specific
# build steps in parallel to avoid encountering an installation race condition.
build.init: $(UP) $(CROSSPLANE_CLI) check-terraform-version

# ====================================================================================
# Setup Terraform for fetching provider schema
TERRAFORM := $(TOOLS_HOST_DIR)/terraform-$(TERRAFORM_VERSION)
TERRAFORM_WORKDIR := $(WORK_DIR)/terraform
TERRAFORM_PROVIDER_SCHEMA := config/schema.json

check-terraform-version:
ifneq ($(TERRAFORM_VERSION_VALID),1)
	$(error invalid TERRAFORM_VERSION $(TERRAFORM_VERSION), must be less than 1.6.0 since that version introduced a not permitted BSL license))
endif

$(TERRAFORM): check-terraform-version
	@$(INFO) installing terraform $(HOSTOS)-$(HOSTARCH)
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-terraform
	@curl -fsSL https://releases.hashicorp.com/terraform/$(TERRAFORM_VERSION)/terraform_$(TERRAFORM_VERSION)_$(SAFEHOST_PLATFORM).zip -o $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip
	@unzip $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip -d $(TOOLS_HOST_DIR)/tmp-terraform
	@mv $(TOOLS_HOST_DIR)/tmp-terraform/terraform $(TERRAFORM)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-terraform
	@$(OK) installing terraform $(HOSTOS)-$(HOSTARCH)

$(TERRAFORM_PROVIDER_SCHEMA): $(TERRAFORM)
	@$(INFO) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)
	@mkdir -p $(TERRAFORM_WORKDIR)
	@echo '{"terraform":[{"required_providers":[{"provider":{"source":"'"$(TERRAFORM_PROVIDER_SOURCE)"'","version":"'"$(TERRAFORM_PROVIDER_VERSION)"'"}}],"required_version":"'"$(TERRAFORM_VERSION)"'"}]}' > $(TERRAFORM_WORKDIR)/main.tf.json
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) init > $(TERRAFORM_WORKDIR)/terraform-logs.txt 2>&1
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) providers schema -json=true > $(TERRAFORM_PROVIDER_SCHEMA) 2>> $(TERRAFORM_WORKDIR)/terraform-logs.txt
	@$(OK) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)

pull-docs:
	@if [ ! -d "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" ]; then \
  		mkdir -p "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" && \
		git clone -c advice.detachedHead=false --depth 1 --filter=blob:none --branch "v$(TERRAFORM_PROVIDER_VERSION)" --sparse "$(TERRAFORM_PROVIDER_REPO)" "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)"; \
	fi
	@git -C "$(WORK_DIR)/$(TERRAFORM_PROVIDER_SOURCE)" sparse-checkout set "$(TERRAFORM_DOCS_PATH)"

generate.init: $(TERRAFORM_PROVIDER_SCHEMA) pull-docs

.PHONY: $(TERRAFORM_PROVIDER_SCHEMA) pull-docs check-terraform-version
# ====================================================================================
# Targets

# NOTE: the build submodule currently overrides XDG_CACHE_HOME in order to
# force the Helm 3 to use the .work/helm directory. This causes Go on Linux
# machines to use that directory as the build cache as well. We should adjust
# this behavior in the build submodule because it is also causing Linux users
# to duplicate their build cache, but for now we just make it easier to identify
# its location in CI so that we cache between builds.
go.cachedir:
	@go env GOCACHE

go.mod.cachedir:
	@go env GOMODCACHE

# Generate a coverage report for cobertura applying exclusions on
# - generated file
cobertura:
	@cat $(GO_TEST_OUTPUT)/coverage.txt | \
		grep -v zz_ | \
		$(GOCOVER_COBERTURA) > $(GO_TEST_OUTPUT)/cobertura-coverage.xml

# Update the submodules, such as the common build scripts.
submodules:
	@git submodule sync
	@git submodule update --init --recursive

# This is for running out-of-cluster locally, and is for convenience. Running
# this make target will print out the command which was used. For more control,
# try running the binary directly with different arguments.
run: go.build
	@$(INFO) Running Crossplane locally out-of-cluster . . .
	@# To see other arguments that can be provided, run the command with --help instead
	$(GO_OUT_DIR)/provider --debug

# ====================================================================================
# End to End Testing
CROSSPLANE_NAMESPACE = crossplane-system
-include build/makelib/local.xpkg.mk
-include build/makelib/controlplane.mk

# This target requires the following environment variables to be set:
# - UPTEST_EXAMPLE_LIST, a comma-separated list of examples to test
#   To ensure the proper functioning of the end-to-end test resource pre-deletion hook, it is crucial to arrange your resources appropriately.
#   You can check the basic implementation here: https://github.com/crossplane/uptest/blob/main/internal/templates/03-delete.yaml.tmpl.
# - UPTEST_CLOUD_CREDENTIALS (optional), multiple sets of AWS IAM User credentials specified as key=value pairs.
#   The support keys are currently `DEFAULT` and `PEER`. So, an example for the value of this env. variable is:
#   DEFAULT='[default]
#   aws_access_key_id = REDACTED
#   aws_secret_access_key = REDACTED'
#   PEER='[default]
#   aws_access_key_id = REDACTED
#   aws_secret_access_key = REDACTED'
#   The associated `ProviderConfig`s will be named as `default` and `peer`.
# - UPTEST_DATASOURCE_PATH (optional), please see https://github.com/crossplane/uptest#injecting-dynamic-values-and-datasource
uptest: $(UPTEST) $(KUBECTL) $(CHAINSAW) $(CROSSPLANE_CLI)
	@$(INFO) running automated tests
	@KUBECTL=$(KUBECTL) CHAINSAW=$(CHAINSAW) CROSSPLANE_CLI=$(CROSSPLANE_CLI) CROSSPLANE_NAMESPACE=$(CROSSPLANE_NAMESPACE) $(UPTEST) e2e "${UPTEST_EXAMPLE_LIST}" --data-source="${UPTEST_DATASOURCE_PATH}" --setup-script=cluster/test/setup.sh --default-conditions="Test" || $(FAIL)
	@$(OK) running automated tests

local-deploy: build controlplane.up local.xpkg.deploy.provider.$(PROJECT_NAME)
	@$(INFO) running locally built provider
	@$(KUBECTL) wait provider.pkg $(PROJECT_NAME) --for condition=Healthy --timeout 5m
	@$(KUBECTL) -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m
	@$(OK) running locally built provider

e2e: local-deploy uptest

crddiff: $(UPTEST)
	@$(INFO) Checking breaking CRD schema changes
	@for crd in $${MODIFIED_CRD_LIST}; do \
		if ! git cat-file -e "$${GITHUB_BASE_REF}:$${crd}" 2>/dev/null; then \
			echo "CRD $${crd} does not exist in the $${GITHUB_BASE_REF} branch. Skipping..." ; \
			continue ; \
		fi ; \
		echo "Checking $${crd} for breaking API changes..." ; \
		changes_detected=$$(go run github.com/crossplane/uptest/cmd/crddiff@$(CRDDIFF_VERSION) revision --enable-upjet-extensions <(git cat-file -p "$${GITHUB_BASE_REF}:$${crd}") "$${crd}" 2>&1) ; \
		if [[ $$? != 0 ]] ; then \
			printf "\033[31m"; echo "Breaking change detected!"; printf "\033[0m" ; \
			echo "$${changes_detected}" ; \
			echo ; \
		fi ; \
	done
	@$(OK) Checking breaking CRD schema changes

schema-version-diff:
	@$(INFO) Checking for native state schema version changes
	@export PREV_PROVIDER_VERSION=$$(git cat-file -p "${GITHUB_BASE_REF}:Makefile" | sed -nr 's/^export[[:space:]]*TERRAFORM_PROVIDER_VERSION[[:space:]]*:=[[:space:]]*(.+)/\1/p'); \
	echo Detected previous Terraform provider version: $${PREV_PROVIDER_VERSION}; \
	echo Current Terraform provider version: $${TERRAFORM_PROVIDER_VERSION}; \
	mkdir -p $(WORK_DIR); \
	git cat-file -p "$${GITHUB_BASE_REF}:config/schema.json" > "$(WORK_DIR)/schema.json.$${PREV_PROVIDER_VERSION}"; \
	./scripts/version_diff.py config/generated.lst "$(WORK_DIR)/schema.json.$${PREV_PROVIDER_VERSION}" config/schema.json
	@$(OK) Checking for native state schema version changes

# docs (E4-S01): render the generated CRD API reference from the Go API types
# under apis/** into docs/api/ using elastic/crd-ref-docs. Runs via
# `go run @pinned-version` so no separate install step is needed, consistent
# with the other tool targets. Deterministic: re-running produces no diff, so a
# CI sync-check can diff docs/api/ later. Config lives in .crd-ref-docs.yaml.
CRD_REF_DOCS_VERSION ?= v0.2.0
docs:
	@$(INFO) generating CRD API reference
	@mkdir -p docs/api
	go run github.com/elastic/crd-ref-docs@$(CRD_REF_DOCS_VERSION) \
		--source-path=apis \
		--config=.crd-ref-docs.yaml \
		--renderer=markdown \
		--output-path=docs/api/
	@$(OK) generating CRD API reference

.PHONY: docs cobertura submodules fallthrough run crds.clean

# ====================================================================================
# Dev Quality (local) targets
#
# Thin local-only quality mechanisms ported from the sibling kollect repo
# (E2-S06/S07/S08/S10). CI wiring is intentionally out of scope here. All four
# targets are hermetic-ish `go run @pinned-version` invocations plus git checks
# so they need no extra tooling install step.

# vuln (E2-S06): report known vulnerabilities in dependencies / std lib via
# govulncheck. Informational: a non-zero exit means real findings to surface,
# not a broken target.
vuln:
	@$(INFO) running govulncheck
	go run golang.org/x/vuln/cmd/govulncheck@v1.1.4 ./...
	@$(OK) running govulncheck

# test.race (E2-S08): run the hand-authored unit tests under the race detector.
# The race detector requires cgo, so force CGO_ENABLED=1. Generated controller
# packages under internal/... carry no _test.go files, so they compile and
# report "no test files" (they do not need envtest at this scope).
test.race:
	@$(INFO) running tests with the race detector
	CGO_ENABLED=1 go test -race -count=1 ./config/... ./internal/...
	@$(OK) running tests with the race detector

# tidy-check (E2-S10): fail if `go mod tidy` would change go.mod/go.sum.
tidy-check:
	@$(INFO) checking go module hygiene
	go mod tidy
	git diff --exit-code go.mod go.sum
	@$(OK) checking go module hygiene

# arch-lint (E2-S07): enforce the generation boundary. Asserts hand-authored
# config/ and internal/clients do not import the generated internal/controller
# tree. Config lives in .go-arch-lint.yml.
arch-lint:
	@$(INFO) checking architecture boundaries
	go run github.com/fe3dback/go-arch-lint@v1.15.0 check --arch-file .go-arch-lint.yml
	@$(OK) checking architecture boundaries

# coverage (E2-S03): unit-test coverage scoped to the hand-authored ./config/...
# surface. Enforces a floor (COVERAGE_MIN, default 70) and emits cover.out as an
# artifact for E5 (codecov). config coverage is ~100% today; the floor guards
# against regression, it is not the target.
COVERAGE_MIN ?= 70
coverage:
	@$(INFO) running config coverage
	go test -covermode=count -coverpkg=./config/... -coverprofile=cover.out ./config/...
	@total=$$(go tool cover -func=cover.out | awk '/^total:/ {gsub(/%/,"",$$3); print $$3}'); \
	echo "total config coverage: $$total% / floor $(COVERAGE_MIN)%"; \
	awk -v t=$$total -v m=$(COVERAGE_MIN) 'BEGIN { exit (t+0 < m+0) }' || { echo "coverage below floor"; exit 1; }
	@$(OK) config coverage floor satisfied

.PHONY: vuln test.race tidy-check arch-lint coverage

# ====================================================================================
# Special Targets

define CROSSPLANE_MAKE_HELP
Crossplane Targets:
    docs                  Generate the CRD API reference into docs/api/ via crd-ref-docs.
    cobertura             Generate a coverage report for cobertura applying exclusions on generated files.
    submodules            Update the submodules, such as the common build scripts.
    run                   Run crossplane locally, out-of-cluster. Useful for development.

endef
# The reason CROSSPLANE_MAKE_HELP is used instead of CROSSPLANE_HELP is because the crossplane
# binary will try to use CROSSPLANE_HELP if it is set, and this is for something different.
export CROSSPLANE_MAKE_HELP

crossplane.help:
	@echo "$$CROSSPLANE_MAKE_HELP"

help-special: crossplane.help

.PHONY: crossplane.help help-special

# TODO(negz): Update CI to use these targets.
vendor: modules.download
vendor.check: modules.check
