package awss3


// Experimental.
type TfBucketAnalyticsConfiguration_S3BucketDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#bucket_arn TfBucketAnalyticsConfiguration#bucket_arn}.
	// Experimental.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#bucket_account_id TfBucketAnalyticsConfiguration#bucket_account_id}.
	// Experimental.
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#format TfBucketAnalyticsConfiguration#format}.
	// Experimental.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_analytics_configuration#prefix TfBucketAnalyticsConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

