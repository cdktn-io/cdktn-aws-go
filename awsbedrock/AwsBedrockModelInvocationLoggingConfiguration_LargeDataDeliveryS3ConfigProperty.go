package awsbedrock


// Experimental.
type AwsBedrockModelInvocationLoggingConfiguration_LargeDataDeliveryS3ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#bucket_name AwsBedrockModelInvocationLoggingConfiguration#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_model_invocation_logging_configuration#key_prefix AwsBedrockModelInvocationLoggingConfiguration#key_prefix}.
	// Experimental.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

