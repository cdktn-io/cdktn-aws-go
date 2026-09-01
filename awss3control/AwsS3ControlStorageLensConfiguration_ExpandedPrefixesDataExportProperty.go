package awss3control


// Experimental.
type AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportProperty struct {
	// s3_bucket_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#s3_bucket_destination AwsS3ControlStorageLensConfiguration#s3_bucket_destination}
	// Experimental.
	S3BucketDestination *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationProperty `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
	// storage_lens_table_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#storage_lens_table_destination AwsS3ControlStorageLensConfiguration#storage_lens_table_destination}
	// Experimental.
	StorageLensTableDestination *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationProperty `field:"optional" json:"storageLensTableDestination" yaml:"storageLensTableDestination"`
}

