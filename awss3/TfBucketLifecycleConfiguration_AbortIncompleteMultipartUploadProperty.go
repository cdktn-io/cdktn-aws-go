package awss3


// Experimental.
type TfBucketLifecycleConfiguration_AbortIncompleteMultipartUploadProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_lifecycle_configuration#days_after_initiation TfBucketLifecycleConfiguration#days_after_initiation}.
	// Experimental.
	DaysAfterInitiation *float64 `field:"optional" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

