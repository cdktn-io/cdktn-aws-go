package s3


// Experimental.
type AwsBucket_ServerSideEncryptionConfigurationRuleProperty struct {
	// apply_server_side_encryption_by_default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#apply_server_side_encryption_by_default AwsBucket#apply_server_side_encryption_by_default}
	// Experimental.
	ApplyServerSideEncryptionByDefault *AwsBucket_ApplyServerSideEncryptionByDefaultProperty `field:"required" json:"applyServerSideEncryptionByDefault" yaml:"applyServerSideEncryptionByDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#bucket_key_enabled AwsBucket#bucket_key_enabled}.
	// Experimental.
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
}

