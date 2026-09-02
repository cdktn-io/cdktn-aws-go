package awss3control


// Experimental.
type TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty struct {
	// s3_bucket_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#s3_bucket_destination TfStorageLensConfiguration#s3_bucket_destination}
	// Experimental.
	S3BucketDestination *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationProperty `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
	// storage_lens_table_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#storage_lens_table_destination TfStorageLensConfiguration#storage_lens_table_destination}
	// Experimental.
	StorageLensTableDestination *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationProperty `field:"optional" json:"storageLensTableDestination" yaml:"storageLensTableDestination"`
}

