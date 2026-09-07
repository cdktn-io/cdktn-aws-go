package s3control


// Experimental.
type AwsBucketLifecycleConfiguration_AbortIncompleteMultipartUploadProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_bucket_lifecycle_configuration#days_after_initiation AwsBucketLifecycleConfiguration#days_after_initiation}.
	// Experimental.
	DaysAfterInitiation *float64 `field:"required" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

