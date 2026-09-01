package awss3control


// Experimental.
type AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionProperty struct {
	// sse_kms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#sse_kms AwsS3ControlStorageLensConfiguration#sse_kms}
	// Experimental.
	SseKms *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionSseKmsProperty `field:"optional" json:"sseKms" yaml:"sseKms"`
	// sse_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#sse_s3 AwsS3ControlStorageLensConfiguration#sse_s3}
	// Experimental.
	SseS3 interface{} `field:"optional" json:"sseS3" yaml:"sseS3"`
}

