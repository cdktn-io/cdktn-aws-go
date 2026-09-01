package awss3


// Experimental.
type AwsS3BucketLogging_TargetObjectKeyFormatProperty struct {
	// partitioned_prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#partitioned_prefix AwsS3BucketLogging#partitioned_prefix}
	// Experimental.
	PartitionedPrefix *AwsS3BucketLogging_PartitionedPrefixProperty `field:"optional" json:"partitionedPrefix" yaml:"partitionedPrefix"`
	// simple_prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#simple_prefix AwsS3BucketLogging#simple_prefix}
	// Experimental.
	SimplePrefix *AwsS3BucketLogging_SimplePrefixProperty `field:"optional" json:"simplePrefix" yaml:"simplePrefix"`
}

