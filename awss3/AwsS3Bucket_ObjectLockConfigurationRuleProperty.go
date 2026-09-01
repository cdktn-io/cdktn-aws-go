package awss3


// Experimental.
type AwsS3Bucket_ObjectLockConfigurationRuleProperty struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#default_retention AwsS3Bucket#default_retention}
	// Experimental.
	DefaultRetention *AwsS3Bucket_DefaultRetentionProperty `field:"required" json:"defaultRetention" yaml:"defaultRetention"`
}

