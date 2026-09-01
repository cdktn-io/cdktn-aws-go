package awss3


// Experimental.
type AwsS3BucketObjectLockConfiguration_DefaultRetentionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#days AwsS3BucketObjectLockConfiguration#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#mode AwsS3BucketObjectLockConfiguration#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#years AwsS3BucketObjectLockConfiguration#years}.
	// Experimental.
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

