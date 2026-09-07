package sagemakerai


// Experimental.
type AwsHyperParameterTuningJob_TrainingJobDefinitionsStoppingConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_pending_time_in_seconds AwsHyperParameterTuningJob#max_pending_time_in_seconds}.
	// Experimental.
	MaxPendingTimeInSeconds *float64 `field:"optional" json:"maxPendingTimeInSeconds" yaml:"maxPendingTimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_runtime_in_seconds AwsHyperParameterTuningJob#max_runtime_in_seconds}.
	// Experimental.
	MaxRuntimeInSeconds *float64 `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_hyper_parameter_tuning_job#max_wait_time_in_seconds AwsHyperParameterTuningJob#max_wait_time_in_seconds}.
	// Experimental.
	MaxWaitTimeInSeconds *float64 `field:"optional" json:"maxWaitTimeInSeconds" yaml:"maxWaitTimeInSeconds"`
}

