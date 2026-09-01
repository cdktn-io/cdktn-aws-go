package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_InputDataConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#channel_name AwsSagemakerAlgorithm#channel_name}.
	// Experimental.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#compression_type AwsSagemakerAlgorithm#compression_type}.
	// Experimental.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#content_type AwsSagemakerAlgorithm#content_type}.
	// Experimental.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#data_source AwsSagemakerAlgorithm#data_source}
	// Experimental.
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#input_mode AwsSagemakerAlgorithm#input_mode}.
	// Experimental.
	InputMode *string `field:"optional" json:"inputMode" yaml:"inputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#record_wrapper_type AwsSagemakerAlgorithm#record_wrapper_type}.
	// Experimental.
	RecordWrapperType *string `field:"optional" json:"recordWrapperType" yaml:"recordWrapperType"`
	// shuffle_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#shuffle_config AwsSagemakerAlgorithm#shuffle_config}
	// Experimental.
	ShuffleConfig interface{} `field:"optional" json:"shuffleConfig" yaml:"shuffleConfig"`
}

