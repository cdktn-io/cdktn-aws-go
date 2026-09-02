package awss3


// Experimental.
type TfBucket_ObjectLockConfigurationRuleProperty struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#default_retention TfBucket#default_retention}
	// Experimental.
	DefaultRetention *TfBucket_DefaultRetentionProperty `field:"required" json:"defaultRetention" yaml:"defaultRetention"`
}

