package awssagemakerai


// Experimental.
type AwsSagemakerEndpointConfiguration_OutputConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#s3_output_path AwsSagemakerEndpointConfiguration#s3_output_path}.
	// Experimental.
	S3OutputPath *string `field:"required" json:"s3OutputPath" yaml:"s3OutputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#kms_key_id AwsSagemakerEndpointConfiguration#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#notification_config AwsSagemakerEndpointConfiguration#notification_config}
	// Experimental.
	NotificationConfig *AwsSagemakerEndpointConfiguration_NotificationConfigProperty `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#s3_failure_path AwsSagemakerEndpointConfiguration#s3_failure_path}.
	// Experimental.
	S3FailurePath *string `field:"optional" json:"s3FailurePath" yaml:"s3FailurePath"`
}

