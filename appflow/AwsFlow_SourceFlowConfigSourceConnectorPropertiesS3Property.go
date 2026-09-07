package appflow


// Experimental.
type AwsFlow_SourceFlowConfigSourceConnectorPropertiesS3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_name AwsFlow#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_prefix AwsFlow#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"required" json:"bucketPrefix" yaml:"bucketPrefix"`
	// s3_input_format_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3_input_format_config AwsFlow#s3_input_format_config}
	// Experimental.
	S3InputFormatConfig *AwsFlow_S3InputFormatConfigProperty `field:"optional" json:"s3InputFormatConfig" yaml:"s3InputFormatConfig"`
}

