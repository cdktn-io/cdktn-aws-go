package awss3


// Experimental.
type TfBucketMetadataConfiguration_RecordExpirationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#expiration TfBucketMetadataConfiguration#expiration}.
	// Experimental.
	Expiration *string `field:"required" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#days TfBucketMetadataConfiguration#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

