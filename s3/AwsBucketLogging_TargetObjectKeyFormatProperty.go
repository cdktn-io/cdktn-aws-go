package s3


// Experimental.
type AwsBucketLogging_TargetObjectKeyFormatProperty struct {
	// partitioned_prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#partitioned_prefix AwsBucketLogging#partitioned_prefix}
	// Experimental.
	PartitionedPrefix *AwsBucketLogging_PartitionedPrefixProperty `field:"optional" json:"partitionedPrefix" yaml:"partitionedPrefix"`
	// simple_prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#simple_prefix AwsBucketLogging#simple_prefix}
	// Experimental.
	SimplePrefix *AwsBucketLogging_SimplePrefixProperty `field:"optional" json:"simplePrefix" yaml:"simplePrefix"`
}

