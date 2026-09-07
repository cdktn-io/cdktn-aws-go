package s3control


// Experimental.
type AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#enabled AwsStorageLensConfiguration#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#encryption AwsStorageLensConfiguration#encryption}
	// Experimental.
	Encryption *AwsStorageLensConfiguration_StorageLensConfigurationDataExportStorageLensTableDestinationEncryptionProperty `field:"optional" json:"encryption" yaml:"encryption"`
}

