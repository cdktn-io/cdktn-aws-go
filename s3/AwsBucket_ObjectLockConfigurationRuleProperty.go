package s3


// Experimental.
type AwsBucket_ObjectLockConfigurationRuleProperty struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#default_retention AwsBucket#default_retention}
	// Experimental.
	DefaultRetention *AwsBucket_DefaultRetentionProperty `field:"required" json:"defaultRetention" yaml:"defaultRetention"`
}

