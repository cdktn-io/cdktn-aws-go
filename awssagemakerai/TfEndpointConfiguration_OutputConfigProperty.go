package awssagemakerai


// Experimental.
type TfEndpointConfiguration_OutputConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#s3_output_path TfEndpointConfiguration#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#kms_key_id TfEndpointConfiguration#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#notification_config TfEndpointConfiguration#notification_config}
	// Experimental.
	NotificationConfig *TfEndpointConfiguration_NotificationConfigProperty `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#s3_failure_path TfEndpointConfiguration#s3_failure_path}.
	// Experimental.
	S3FailurePath *string `field:"optional" json:"s3FailurePath" yaml:"s3FailurePath"`
}

