package awss3


// Experimental.
type AwsS3BucketMetadataConfiguration_JournalTableConfigurationProperty struct {
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#encryption_configuration AwsS3BucketMetadataConfiguration#encryption_configuration}
	// Experimental.
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// record_expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#record_expiration AwsS3BucketMetadataConfiguration#record_expiration}
	// Experimental.
	RecordExpiration interface{} `field:"optional" json:"recordExpiration" yaml:"recordExpiration"`
}

