package awss3


// Experimental.
type TfBucketAnalyticsConfiguration_DestinationProperty struct {
	// s3_bucket_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#s3_bucket_destination TfBucketAnalyticsConfiguration#s3_bucket_destination}
	// Experimental.
	S3BucketDestination *TfBucketAnalyticsConfiguration_S3BucketDestinationProperty `field:"required" json:"s3BucketDestination" yaml:"s3BucketDestination"`
}

