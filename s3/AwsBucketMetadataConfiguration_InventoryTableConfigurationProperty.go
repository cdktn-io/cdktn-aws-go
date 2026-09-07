package s3


// Experimental.
type AwsBucketMetadataConfiguration_InventoryTableConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#configuration_state AwsBucketMetadataConfiguration#configuration_state}.
	// Experimental.
	ConfigurationState *string `field:"required" json:"configurationState" yaml:"configurationState"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#encryption_configuration AwsBucketMetadataConfiguration#encryption_configuration}
	// Experimental.
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

