package awssagemakerai


// Experimental.
type AwsSagemakerEndpoint_RollingUpdatePolicyProperty struct {
	// maximum_batch_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#maximum_batch_size AwsSagemakerEndpoint#maximum_batch_size}
	// Experimental.
	MaximumBatchSize *AwsSagemakerEndpoint_MaximumBatchSizeProperty `field:"required" json:"maximumBatchSize" yaml:"maximumBatchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#wait_interval_in_seconds AwsSagemakerEndpoint#wait_interval_in_seconds}.
	// Experimental.
	WaitIntervalInSeconds *float64 `field:"required" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#maximum_execution_timeout_in_seconds AwsSagemakerEndpoint#maximum_execution_timeout_in_seconds}.
	// Experimental.
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// rollback_maximum_batch_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#rollback_maximum_batch_size AwsSagemakerEndpoint#rollback_maximum_batch_size}
	// Experimental.
	RollbackMaximumBatchSize *AwsSagemakerEndpoint_RollbackMaximumBatchSizeProperty `field:"optional" json:"rollbackMaximumBatchSize" yaml:"rollbackMaximumBatchSize"`
}

