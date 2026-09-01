package awss3


// Experimental.
type AwsS3BucketReplicationConfiguration_ReplicationTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#status AwsS3BucketReplicationConfiguration#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#time AwsS3BucketReplicationConfiguration#time}
	// Experimental.
	Time *AwsS3BucketReplicationConfiguration_TimeProperty `field:"required" json:"time" yaml:"time"`
}

