package awsappflow


// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_name TfFlow#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#bucket_prefix TfFlow#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// s3_output_format_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#s3_output_format_config TfFlow#s3_output_format_config}
	// Experimental.
	S3OutputFormatConfig *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigProperty `field:"optional" json:"s3OutputFormatConfig" yaml:"s3OutputFormatConfig"`
}

