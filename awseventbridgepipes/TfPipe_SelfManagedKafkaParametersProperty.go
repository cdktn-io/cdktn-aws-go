package awseventbridgepipes


// Experimental.
type TfPipe_SelfManagedKafkaParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#topic_name TfPipe#topic_name}.
	// Experimental.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#additional_bootstrap_servers TfPipe#additional_bootstrap_servers}.
	// Experimental.
	AdditionalBootstrapServers *[]*string `field:"optional" json:"additionalBootstrapServers" yaml:"additionalBootstrapServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#batch_size TfPipe#batch_size}.
	// Experimental.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#consumer_group_id TfPipe#consumer_group_id}.
	// Experimental.
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#credentials TfPipe#credentials}
	// Experimental.
	Credentials *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty `field:"optional" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#maximum_batching_window_in_seconds TfPipe#maximum_batching_window_in_seconds}.
	// Experimental.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#server_root_ca_certificate TfPipe#server_root_ca_certificate}.
	// Experimental.
	ServerRootCaCertificate *string `field:"optional" json:"serverRootCaCertificate" yaml:"serverRootCaCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#starting_position TfPipe#starting_position}.
	// Experimental.
	StartingPosition *string `field:"optional" json:"startingPosition" yaml:"startingPosition"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#vpc TfPipe#vpc}
	// Experimental.
	Vpc *TfPipe_VpcProperty `field:"optional" json:"vpc" yaml:"vpc"`
}

