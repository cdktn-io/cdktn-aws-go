package s3


// Experimental.
type AwsBucketMetadataConfiguration_RecordExpirationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#expiration AwsBucketMetadataConfiguration#expiration}.
	// Experimental.
	Expiration *string `field:"required" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_metadata_configuration#days AwsBucketMetadataConfiguration#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

