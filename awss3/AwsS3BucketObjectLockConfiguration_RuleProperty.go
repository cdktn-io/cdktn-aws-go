package awss3


// Experimental.
type AwsS3BucketObjectLockConfiguration_RuleProperty struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#default_retention AwsS3BucketObjectLockConfiguration#default_retention}
	// Experimental.
	DefaultRetention *AwsS3BucketObjectLockConfiguration_DefaultRetentionProperty `field:"required" json:"defaultRetention" yaml:"defaultRetention"`
}

