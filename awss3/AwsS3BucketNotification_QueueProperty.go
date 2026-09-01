package awss3


// Experimental.
type AwsS3BucketNotification_QueueProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#events AwsS3BucketNotification#events}.
	// Experimental.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#queue_arn AwsS3BucketNotification#queue_arn}.
	// Experimental.
	QueueArn *string `field:"required" json:"queueArn" yaml:"queueArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#filter_prefix AwsS3BucketNotification#filter_prefix}.
	// Experimental.
	FilterPrefix *string `field:"optional" json:"filterPrefix" yaml:"filterPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#filter_suffix AwsS3BucketNotification#filter_suffix}.
	// Experimental.
	FilterSuffix *string `field:"optional" json:"filterSuffix" yaml:"filterSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_notification#id AwsS3BucketNotification#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

