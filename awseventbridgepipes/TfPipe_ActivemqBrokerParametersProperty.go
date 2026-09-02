package awseventbridgepipes


// Experimental.
type TfPipe_ActivemqBrokerParametersProperty struct {
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#credentials TfPipe#credentials}
	// Experimental.
	Credentials *TfPipe_SourceParametersActivemqBrokerParametersCredentialsProperty `field:"required" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#queue_name TfPipe#queue_name}.
	// Experimental.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#batch_size TfPipe#batch_size}.
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#maximum_batching_window_in_seconds TfPipe#maximum_batching_window_in_seconds}.
	// Experimental.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
}

