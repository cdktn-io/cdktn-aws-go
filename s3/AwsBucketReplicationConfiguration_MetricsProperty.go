package s3


// Experimental.
type AwsBucketReplicationConfiguration_MetricsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#status AwsBucketReplicationConfiguration#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// event_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#event_threshold AwsBucketReplicationConfiguration#event_threshold}
	// Experimental.
	EventThreshold *AwsBucketReplicationConfiguration_EventThresholdProperty `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
}

