package awss3


// Experimental.
type AwsS3BucketReplicationConfiguration_FilterProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#and AwsS3BucketReplicationConfiguration#and}
	// Experimental.
	And *AwsS3BucketReplicationConfiguration_AndProperty `field:"optional" json:"and" yaml:"and"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#prefix AwsS3BucketReplicationConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#tag AwsS3BucketReplicationConfiguration#tag}
	// Experimental.
	Tag *AwsS3BucketReplicationConfiguration_TagProperty `field:"optional" json:"tag" yaml:"tag"`
}

