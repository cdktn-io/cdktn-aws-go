package awss3


// Experimental.
type TfBucketReplicationConfiguration_FilterProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#and TfBucketReplicationConfiguration#and}
	// Experimental.
	And *TfBucketReplicationConfiguration_AndProperty `field:"optional" json:"and" yaml:"and"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#prefix TfBucketReplicationConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#tag TfBucketReplicationConfiguration#tag}
	// Experimental.
	Tag *TfBucketReplicationConfiguration_TagProperty `field:"optional" json:"tag" yaml:"tag"`
}

