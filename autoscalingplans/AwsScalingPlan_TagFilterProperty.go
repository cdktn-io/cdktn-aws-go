package autoscalingplans


// Experimental.
type AwsScalingPlan_TagFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#key AwsScalingPlan#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#values AwsScalingPlan#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

