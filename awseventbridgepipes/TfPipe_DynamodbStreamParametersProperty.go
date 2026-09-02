package awseventbridgepipes


// Experimental.
type TfPipe_DynamodbStreamParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#starting_position TfPipe#starting_position}.
	// Experimental.
	StartingPosition *string `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#batch_size TfPipe#batch_size}.
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#dead_letter_config TfPipe#dead_letter_config}
	// Experimental.
	DeadLetterConfig *TfPipe_SourceParametersDynamodbStreamParametersDeadLetterConfigProperty `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#maximum_batching_window_in_seconds TfPipe#maximum_batching_window_in_seconds}.
	// Experimental.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#maximum_record_age_in_seconds TfPipe#maximum_record_age_in_seconds}.
	// Experimental.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#maximum_retry_attempts TfPipe#maximum_retry_attempts}.
	// Experimental.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#on_partial_batch_item_failure TfPipe#on_partial_batch_item_failure}.
	// Experimental.
	OnPartialBatchItemFailure *string `field:"optional" json:"onPartialBatchItemFailure" yaml:"onPartialBatchItemFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#parallelization_factor TfPipe#parallelization_factor}.
	// Experimental.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
}

