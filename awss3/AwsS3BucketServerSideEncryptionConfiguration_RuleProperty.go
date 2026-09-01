package awss3


// Experimental.
type AwsS3BucketServerSideEncryptionConfiguration_RuleProperty struct {
	// apply_server_side_encryption_by_default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_server_side_encryption_configuration#apply_server_side_encryption_by_default AwsS3BucketServerSideEncryptionConfiguration#apply_server_side_encryption_by_default}
	// Experimental.
	ApplyServerSideEncryptionByDefault *AwsS3BucketServerSideEncryptionConfiguration_ApplyServerSideEncryptionByDefaultProperty `field:"optional" json:"applyServerSideEncryptionByDefault" yaml:"applyServerSideEncryptionByDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_server_side_encryption_configuration#blocked_encryption_types AwsS3BucketServerSideEncryptionConfiguration#blocked_encryption_types}.
	// Experimental.
	BlockedEncryptionTypes *[]*string `field:"optional" json:"blockedEncryptionTypes" yaml:"blockedEncryptionTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_server_side_encryption_configuration#bucket_key_enabled AwsS3BucketServerSideEncryptionConfiguration#bucket_key_enabled}.
	// Experimental.
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
}

