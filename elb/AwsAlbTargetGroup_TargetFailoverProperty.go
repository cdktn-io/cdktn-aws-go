package elb


// Experimental.
type AwsAlbTargetGroup_TargetFailoverProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#on_deregistration AwsAlbTargetGroup#on_deregistration}.
	// Experimental.
	OnDeregistration *string `field:"required" json:"onDeregistration" yaml:"onDeregistration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#on_unhealthy AwsAlbTargetGroup#on_unhealthy}.
	// Experimental.
	OnUnhealthy *string `field:"required" json:"onUnhealthy" yaml:"onUnhealthy"`
}

