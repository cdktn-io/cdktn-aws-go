package sagemakerai


// Experimental.
type AwsFlowDefinition_HumanLoopConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#human_task_ui_arn AwsFlowDefinition#human_task_ui_arn}.
	// Experimental.
	HumanTaskUiArn *string `field:"required" json:"humanTaskUiArn" yaml:"humanTaskUiArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#task_count AwsFlowDefinition#task_count}.
	// Experimental.
	TaskCount *float64 `field:"required" json:"taskCount" yaml:"taskCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#task_description AwsFlowDefinition#task_description}.
	// Experimental.
	TaskDescription *string `field:"required" json:"taskDescription" yaml:"taskDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#task_title AwsFlowDefinition#task_title}.
	// Experimental.
	TaskTitle *string `field:"required" json:"taskTitle" yaml:"taskTitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#workteam_arn AwsFlowDefinition#workteam_arn}.
	// Experimental.
	WorkteamArn *string `field:"required" json:"workteamArn" yaml:"workteamArn"`
	// public_workforce_task_price block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#public_workforce_task_price AwsFlowDefinition#public_workforce_task_price}
	// Experimental.
	PublicWorkforceTaskPrice *AwsFlowDefinition_PublicWorkforceTaskPriceProperty `field:"optional" json:"publicWorkforceTaskPrice" yaml:"publicWorkforceTaskPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#task_availability_lifetime_in_seconds AwsFlowDefinition#task_availability_lifetime_in_seconds}.
	// Experimental.
	TaskAvailabilityLifetimeInSeconds *float64 `field:"optional" json:"taskAvailabilityLifetimeInSeconds" yaml:"taskAvailabilityLifetimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#task_keywords AwsFlowDefinition#task_keywords}.
	// Experimental.
	TaskKeywords *[]*string `field:"optional" json:"taskKeywords" yaml:"taskKeywords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#task_time_limit_in_seconds AwsFlowDefinition#task_time_limit_in_seconds}.
	// Experimental.
	TaskTimeLimitInSeconds *float64 `field:"optional" json:"taskTimeLimitInSeconds" yaml:"taskTimeLimitInSeconds"`
}

