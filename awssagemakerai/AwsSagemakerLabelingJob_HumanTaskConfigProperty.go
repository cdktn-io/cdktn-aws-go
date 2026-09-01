package awssagemakerai


// Experimental.
type AwsSagemakerLabelingJob_HumanTaskConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#number_of_human_workers_per_data_object AwsSagemakerLabelingJob#number_of_human_workers_per_data_object}.
	// Experimental.
	NumberOfHumanWorkersPerDataObject *float64 `field:"required" json:"numberOfHumanWorkersPerDataObject" yaml:"numberOfHumanWorkersPerDataObject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#task_description AwsSagemakerLabelingJob#task_description}.
	// Experimental.
	TaskDescription *string `field:"required" json:"taskDescription" yaml:"taskDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#task_time_limit_in_seconds AwsSagemakerLabelingJob#task_time_limit_in_seconds}.
	// Experimental.
	TaskTimeLimitInSeconds *float64 `field:"required" json:"taskTimeLimitInSeconds" yaml:"taskTimeLimitInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#task_title AwsSagemakerLabelingJob#task_title}.
	// Experimental.
	TaskTitle *string `field:"required" json:"taskTitle" yaml:"taskTitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#workteam_arn AwsSagemakerLabelingJob#workteam_arn}.
	// Experimental.
	WorkteamArn *string `field:"required" json:"workteamArn" yaml:"workteamArn"`
	// annotation_consolidation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#annotation_consolidation_config AwsSagemakerLabelingJob#annotation_consolidation_config}
	// Experimental.
	AnnotationConsolidationConfig interface{} `field:"optional" json:"annotationConsolidationConfig" yaml:"annotationConsolidationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#max_concurrent_task_count AwsSagemakerLabelingJob#max_concurrent_task_count}.
	// Experimental.
	MaxConcurrentTaskCount *float64 `field:"optional" json:"maxConcurrentTaskCount" yaml:"maxConcurrentTaskCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#pre_human_task_lambda_arn AwsSagemakerLabelingJob#pre_human_task_lambda_arn}.
	// Experimental.
	PreHumanTaskLambdaArn *string `field:"optional" json:"preHumanTaskLambdaArn" yaml:"preHumanTaskLambdaArn"`
	// public_workforce_task_price block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#public_workforce_task_price AwsSagemakerLabelingJob#public_workforce_task_price}
	// Experimental.
	PublicWorkforceTaskPrice interface{} `field:"optional" json:"publicWorkforceTaskPrice" yaml:"publicWorkforceTaskPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#task_availability_lifetime_in_seconds AwsSagemakerLabelingJob#task_availability_lifetime_in_seconds}.
	// Experimental.
	TaskAvailabilityLifetimeInSeconds *float64 `field:"optional" json:"taskAvailabilityLifetimeInSeconds" yaml:"taskAvailabilityLifetimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#task_keywords AwsSagemakerLabelingJob#task_keywords}.
	// Experimental.
	TaskKeywords *[]*string `field:"optional" json:"taskKeywords" yaml:"taskKeywords"`
	// ui_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_labeling_job#ui_config AwsSagemakerLabelingJob#ui_config}
	// Experimental.
	UiConfig interface{} `field:"optional" json:"uiConfig" yaml:"uiConfig"`
}

