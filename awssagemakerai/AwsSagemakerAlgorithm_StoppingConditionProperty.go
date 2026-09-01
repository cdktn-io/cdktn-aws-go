package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_StoppingConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_pending_time_in_seconds AwsSagemakerAlgorithm#max_pending_time_in_seconds}.
	// Experimental.
	MaxPendingTimeInSeconds *float64 `field:"optional" json:"maxPendingTimeInSeconds" yaml:"maxPendingTimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_runtime_in_seconds AwsSagemakerAlgorithm#max_runtime_in_seconds}.
	// Experimental.
	MaxRuntimeInSeconds *float64 `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#max_wait_time_in_seconds AwsSagemakerAlgorithm#max_wait_time_in_seconds}.
	// Experimental.
	MaxWaitTimeInSeconds *float64 `field:"optional" json:"maxWaitTimeInSeconds" yaml:"maxWaitTimeInSeconds"`
}

