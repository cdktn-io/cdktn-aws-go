package awss3


// Experimental.
type TfBucketLogging_TargetObjectKeyFormatProperty struct {
	// partitioned_prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#partitioned_prefix TfBucketLogging#partitioned_prefix}
	// Experimental.
	PartitionedPrefix *TfBucketLogging_PartitionedPrefixProperty `field:"optional" json:"partitionedPrefix" yaml:"partitionedPrefix"`
	// simple_prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#simple_prefix TfBucketLogging#simple_prefix}
	// Experimental.
	SimplePrefix *TfBucketLogging_SimplePrefixProperty `field:"optional" json:"simplePrefix" yaml:"simplePrefix"`
}

