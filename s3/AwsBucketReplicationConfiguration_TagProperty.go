package s3


// Experimental.
type AwsBucketReplicationConfiguration_TagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#key AwsBucketReplicationConfiguration#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#value AwsBucketReplicationConfiguration#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

