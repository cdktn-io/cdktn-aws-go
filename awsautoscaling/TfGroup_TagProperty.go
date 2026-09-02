package awsautoscaling


// Experimental.
type TfGroup_TagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#key TfGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#propagate_at_launch TfGroup#propagate_at_launch}.
	// Experimental.
	PropagateAtLaunch interface{} `field:"required" json:"propagateAtLaunch" yaml:"propagateAtLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#value TfGroup#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

