package awsautoscaling


// Experimental.
type TfGroupTag_TagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group_tag#key TfGroupTag#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group_tag#propagate_at_launch TfGroupTag#propagate_at_launch}.
	// Experimental.
	PropagateAtLaunch interface{} `field:"required" json:"propagateAtLaunch" yaml:"propagateAtLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group_tag#value TfGroupTag#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

