package awss3


// Experimental.
type TfBucketReplicationConfiguration_TagProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#key TfBucketReplicationConfiguration#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#value TfBucketReplicationConfiguration#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

