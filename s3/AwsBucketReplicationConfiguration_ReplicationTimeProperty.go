package s3


// Experimental.
type AwsBucketReplicationConfiguration_ReplicationTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#status AwsBucketReplicationConfiguration#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#time AwsBucketReplicationConfiguration#time}
	// Experimental.
	Time *AwsBucketReplicationConfiguration_TimeProperty `field:"required" json:"time" yaml:"time"`
}

