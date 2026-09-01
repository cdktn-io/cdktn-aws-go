package awsappflow


// Experimental.
type AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_name AwsAppflowFlow#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_prefix AwsAppflowFlow#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"required" json:"bucketPrefix" yaml:"bucketPrefix"`
	// s3_input_format_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3_input_format_config AwsAppflowFlow#s3_input_format_config}
	// Experimental.
	S3InputFormatConfig *AwsAppflowFlow_S3InputFormatConfigProperty `field:"optional" json:"s3InputFormatConfig" yaml:"s3InputFormatConfig"`
}

