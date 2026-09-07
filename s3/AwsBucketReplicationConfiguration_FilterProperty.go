package s3


// Experimental.
type AwsBucketReplicationConfiguration_FilterProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#and AwsBucketReplicationConfiguration#and}
	// Experimental.
	And *AwsBucketReplicationConfiguration_AndProperty `field:"optional" json:"and" yaml:"and"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#prefix AwsBucketReplicationConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#tag AwsBucketReplicationConfiguration#tag}
	// Experimental.
	Tag *AwsBucketReplicationConfiguration_TagProperty `field:"optional" json:"tag" yaml:"tag"`
}

