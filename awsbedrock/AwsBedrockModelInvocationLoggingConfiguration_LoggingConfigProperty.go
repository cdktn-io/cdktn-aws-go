package awsbedrock


// Experimental.
type AwsBedrockModelInvocationLoggingConfiguration_LoggingConfigProperty struct {
	// cloudwatch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#cloudwatch_config AwsBedrockModelInvocationLoggingConfiguration#cloudwatch_config}
	// Experimental.
	CloudwatchConfig interface{} `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#embedding_data_delivery_enabled AwsBedrockModelInvocationLoggingConfiguration#embedding_data_delivery_enabled}.
	// Experimental.
	EmbeddingDataDeliveryEnabled interface{} `field:"optional" json:"embeddingDataDeliveryEnabled" yaml:"embeddingDataDeliveryEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#image_data_delivery_enabled AwsBedrockModelInvocationLoggingConfiguration#image_data_delivery_enabled}.
	// Experimental.
	ImageDataDeliveryEnabled interface{} `field:"optional" json:"imageDataDeliveryEnabled" yaml:"imageDataDeliveryEnabled"`
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#s3_config AwsBedrockModelInvocationLoggingConfiguration#s3_config}
	// Experimental.
	S3Config interface{} `field:"optional" json:"s3Config" yaml:"s3Config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#text_data_delivery_enabled AwsBedrockModelInvocationLoggingConfiguration#text_data_delivery_enabled}.
	// Experimental.
	TextDataDeliveryEnabled interface{} `field:"optional" json:"textDataDeliveryEnabled" yaml:"textDataDeliveryEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#video_data_delivery_enabled AwsBedrockModelInvocationLoggingConfiguration#video_data_delivery_enabled}.
	// Experimental.
	VideoDataDeliveryEnabled interface{} `field:"optional" json:"videoDataDeliveryEnabled" yaml:"videoDataDeliveryEnabled"`
}

