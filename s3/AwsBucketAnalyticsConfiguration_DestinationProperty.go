package s3


// Experimental.
type AwsBucketAnalyticsConfiguration_DestinationProperty struct {
	// s3_bucket_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#s3_bucket_destination AwsBucketAnalyticsConfiguration#s3_bucket_destination}
	// Experimental.
	S3BucketDestination *AwsBucketAnalyticsConfiguration_S3BucketDestinationProperty `field:"required" json:"s3BucketDestination" yaml:"s3BucketDestination"`
}

