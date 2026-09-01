package awss3


// Experimental.
type AwsS3BucketReplicationConfiguration_MetricsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#status AwsS3BucketReplicationConfiguration#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// event_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#event_threshold AwsS3BucketReplicationConfiguration#event_threshold}
	// Experimental.
	EventThreshold *AwsS3BucketReplicationConfiguration_EventThresholdProperty `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
}

