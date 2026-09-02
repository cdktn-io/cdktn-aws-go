package awselb


// Experimental.
type TfTargetGroup_TargetFailoverProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#on_deregistration TfTargetGroup#on_deregistration}.
	// Experimental.
	OnDeregistration *string `field:"required" json:"onDeregistration" yaml:"onDeregistration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#on_unhealthy TfTargetGroup#on_unhealthy}.
	// Experimental.
	OnUnhealthy *string `field:"required" json:"onUnhealthy" yaml:"onUnhealthy"`
}

