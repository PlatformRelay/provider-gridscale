#!/usr/bin/env python3

import json
import sys
from pathlib import Path

# usage: version_diff.py <generated resource list> <base JSON schema path> <bumped JSON schema path>
# example usage: version_diff.py config/generated.lst .work/schema.json.3.38.0 config/schema.json

REPO_ROOT = Path(__file__).resolve().parent.parent


class PathOutsideRepoError(ValueError):
    """Raised when a resolved path escapes the repository root."""


def resolve_under_repo(path: str | Path, *, repo_root: Path | None = None) -> Path:
    """Resolve *path* and reject it unless it stays under the repository root."""
    root = (repo_root if repo_root is not None else REPO_ROOT).resolve()
    resolved = Path(path).resolve()
    try:
        resolved.relative_to(root)
    except ValueError as exc:
        raise PathOutsideRepoError(
            f"path escapes repository root: {path!r} -> {resolved}"
        ) from exc
    return resolved


if __name__ == "__main__":
    try:
        resources_path = resolve_under_repo(sys.argv[1])
        base_path = resolve_under_repo(sys.argv[2])
        bumped_path = resolve_under_repo(sys.argv[3])
    except PathOutsideRepoError as err:
        print(err, file=sys.stderr)
        sys.exit(1)

    print(
        f'Reporting schema changes between "{base_path}" as base version '
        f'and "{bumped_path}" as bumped version'
    )
    with open(resources_path) as f:
        resources = json.load(f)
    with open(base_path) as f:
        base = json.load(f)
    with open(bumped_path) as f:
        bump = json.load(f)

    provider_name = None
    for k in base["provider_schemas"]:
        # the first key is the provider name
        provider_name = k
        break
    if provider_name is None:
        print(f"Cannot extract the provider name from the base schema: {base_path}")
        sys.exit(-1)
    base_schemas = base["provider_schemas"][provider_name]["resource_schemas"]
    bumped_schemas = bump["provider_schemas"][provider_name]["resource_schemas"]

    for name in resources:
        try:
            if base_schemas[name]["version"] != bumped_schemas[name]["version"]:
                print(f'{name}:{base_schemas[name]["version"]}-{bumped_schemas[name]["version"]}')
        except KeyError as ke:
            print(f'{name} is not found in schema: {ke}')
            continue
