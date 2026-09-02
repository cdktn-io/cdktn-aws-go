package awsautoscalingplans


// Experimental.
type TfScalingPlan_ApplicationSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#cloudformation_stack_arn TfScalingPlan#cloudformation_stack_arn}.
	// Experimental.
	CloudformationStackArn *string `field:"optional" json:"cloudformationStackArn" yaml:"cloudformationStackArn"`
	// tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscalingplans_scaling_plan#tag_filter TfScalingPlan#tag_filter}
	// Experimental.
	TagFilter interface{} `field:"optional" json:"tagFilter" yaml:"tagFilter"`
}

