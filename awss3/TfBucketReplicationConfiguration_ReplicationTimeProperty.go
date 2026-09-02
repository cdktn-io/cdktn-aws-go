package awss3


// Experimental.
type TfBucketReplicationConfiguration_ReplicationTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#status TfBucketReplicationConfiguration#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#time TfBucketReplicationConfiguration#time}
	// Experimental.
	Time *TfBucketReplicationConfiguration_TimeProperty `field:"required" json:"time" yaml:"time"`
}

