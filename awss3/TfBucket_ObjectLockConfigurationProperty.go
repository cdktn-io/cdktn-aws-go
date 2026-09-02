package awss3


// Experimental.
type TfBucket_ObjectLockConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#object_lock_enabled TfBucket#object_lock_enabled}.
	// Experimental.
	ObjectLockEnabled *string `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#rule TfBucket#rule}
	// Experimental.
	Rule *TfBucket_ObjectLockConfigurationRuleProperty `field:"optional" json:"rule" yaml:"rule"`
}

