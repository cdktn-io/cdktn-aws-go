package awssagemakerai


// Experimental.
type TfDeviceFleet_OutputConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_device_fleet#s3_output_location TfDeviceFleet#s3_output_location}.
	// Experimental.
	S3OutputLocation *string `field:"required" json:"s3OutputLocation" yaml:"s3OutputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_device_fleet#kms_key_id TfDeviceFleet#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

