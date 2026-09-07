package s3


// Experimental.
type AwsBucketLogging_PartitionedPrefixProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#partition_date_source AwsBucketLogging#partition_date_source}.
	// Experimental.
	PartitionDateSource *string `field:"required" json:"partitionDateSource" yaml:"partitionDateSource"`
}

