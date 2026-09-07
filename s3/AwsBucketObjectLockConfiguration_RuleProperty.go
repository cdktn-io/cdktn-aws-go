package s3


// Experimental.
type AwsBucketObjectLockConfiguration_RuleProperty struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#default_retention AwsBucketObjectLockConfiguration#default_retention}
	// Experimental.
	DefaultRetention *AwsBucketObjectLockConfiguration_DefaultRetentionProperty `field:"required" json:"defaultRetention" yaml:"defaultRetention"`
}

