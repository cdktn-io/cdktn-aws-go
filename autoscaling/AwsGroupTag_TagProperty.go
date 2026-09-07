package autoscaling


// Experimental.
type AwsGroupTag_TagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group_tag#key AwsGroupTag#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group_tag#propagate_at_launch AwsGroupTag#propagate_at_launch}.
	// Experimental.
	PropagateAtLaunch interface{} `field:"required" json:"propagateAtLaunch" yaml:"propagateAtLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group_tag#value AwsGroupTag#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

