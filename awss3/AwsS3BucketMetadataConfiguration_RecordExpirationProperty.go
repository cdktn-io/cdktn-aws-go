package awss3


// Experimental.
type AwsS3BucketMetadataConfiguration_RecordExpirationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#expiration AwsS3BucketMetadataConfiguration#expiration}.
	// Experimental.
	Expiration *string `field:"required" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#days AwsS3BucketMetadataConfiguration#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

