package awss3


// Experimental.
type AwsS3Bucket_ServerSideEncryptionConfigurationRuleProperty struct {
	// apply_server_side_encryption_by_default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#apply_server_side_encryption_by_default AwsS3Bucket#apply_server_side_encryption_by_default}
	// Experimental.
	ApplyServerSideEncryptionByDefault *AwsS3Bucket_ApplyServerSideEncryptionByDefaultProperty `field:"required" json:"applyServerSideEncryptionByDefault" yaml:"applyServerSideEncryptionByDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#bucket_key_enabled AwsS3Bucket#bucket_key_enabled}.
	// Experimental.
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
}

