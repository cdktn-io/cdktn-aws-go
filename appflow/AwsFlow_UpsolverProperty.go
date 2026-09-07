package appflow


// Experimental.
type AwsFlow_UpsolverProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_name AwsFlow#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// s3_output_format_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3_output_format_config AwsFlow#s3_output_format_config}
	// Experimental.
	S3OutputFormatConfig *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty `field:"required" json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_prefix AwsFlow#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

