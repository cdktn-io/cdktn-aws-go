package awss3


// Experimental.
type TfBucketReplicationConfiguration_MetricsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#status TfBucketReplicationConfiguration#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// event_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#event_threshold TfBucketReplicationConfiguration#event_threshold}
	// Experimental.
	EventThreshold *TfBucketReplicationConfiguration_EventThresholdProperty `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
}

