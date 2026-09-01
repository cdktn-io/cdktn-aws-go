package awss3


// Experimental.
type AwsS3Bucket_ObjectLockConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#object_lock_enabled AwsS3Bucket#object_lock_enabled}.
	// Experimental.
	ObjectLockEnabled *string `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#rule AwsS3Bucket#rule}
	// Experimental.
	Rule *AwsS3Bucket_ObjectLockConfigurationRuleProperty `field:"optional" json:"rule" yaml:"rule"`
}

