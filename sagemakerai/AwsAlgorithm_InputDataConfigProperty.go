package sagemakerai


// Experimental.
type AwsAlgorithm_InputDataConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#channel_name AwsAlgorithm#channel_name}.
	// Experimental.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#compression_type AwsAlgorithm#compression_type}.
	// Experimental.
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#content_type AwsAlgorithm#content_type}.
	// Experimental.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#data_source AwsAlgorithm#data_source}
	// Experimental.
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#input_mode AwsAlgorithm#input_mode}.
	// Experimental.
	InputMode *string `field:"optional" json:"inputMode" yaml:"inputMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#record_wrapper_type AwsAlgorithm#record_wrapper_type}.
	// Experimental.
	RecordWrapperType *string `field:"optional" json:"recordWrapperType" yaml:"recordWrapperType"`
	// shuffle_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#shuffle_config AwsAlgorithm#shuffle_config}
	// Experimental.
	ShuffleConfig interface{} `field:"optional" json:"shuffleConfig" yaml:"shuffleConfig"`
}

