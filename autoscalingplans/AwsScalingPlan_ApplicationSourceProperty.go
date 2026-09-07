package autoscalingplans


// Experimental.
type AwsScalingPlan_ApplicationSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#cloudformation_stack_arn AwsScalingPlan#cloudformation_stack_arn}.
	// Experimental.
	CloudformationStackArn *string `field:"optional" json:"cloudformationStackArn" yaml:"cloudformationStackArn"`
	// tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#tag_filter AwsScalingPlan#tag_filter}
	// Experimental.
	TagFilter interface{} `field:"optional" json:"tagFilter" yaml:"tagFilter"`
}

