package config

import "testing"

// TestLoadbalancerStatusComputed asserts the LB-1 override: gridscale_loadbalancer
// "status" must be Computed (observed-only), not Optional. Upstream declares it
// Optional+Default which causes a perpetual plan diff; every other resource marks
// its status Computed. If this regresses, status reappears in spec.forProvider and
// the drift returns.
func TestLoadbalancerStatusComputed(t *testing.T) {
	p := GetProvider()

	r, ok := p.Resources["gridscale_loadbalancer"]
	if !ok {
		t.Fatal("gridscale_loadbalancer not present in the configured provider")
	}
	s, ok := r.TerraformResource.Schema["status"]
	if !ok {
		t.Fatal("gridscale_loadbalancer.status missing from Terraform schema")
	}
	if !s.Computed {
		t.Error("gridscale_loadbalancer.status: want Computed=true (LB-1), got false")
	}
	if s.Optional {
		t.Error("gridscale_loadbalancer.status: want Optional=false (LB-1), got true — would reintroduce the perpetual diff")
	}
	if s.Default != nil {
		t.Errorf("gridscale_loadbalancer.status: want Default=nil (LB-1), got %v", s.Default)
	}
}
