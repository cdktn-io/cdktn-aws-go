package s3


// Experimental.
type AwsBucketReplicationConfiguration_EventThresholdProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#minutes AwsBucketReplicationConfiguration#minutes}.
	// Experimental.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

