package awss3


// Experimental.
type TfBucketLifecycleConfiguration_NoncurrentVersionExpirationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#noncurrent_days TfBucketLifecycleConfiguration#noncurrent_days}.
	// Experimental.
	NoncurrentDays *float64 `field:"required" json:"noncurrentDays" yaml:"noncurrentDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#newer_noncurrent_versions TfBucketLifecycleConfiguration#newer_noncurrent_versions}.
	// Experimental.
	NewerNoncurrentVersions *float64 `field:"optional" json:"newerNoncurrentVersions" yaml:"newerNoncurrentVersions"`
}

