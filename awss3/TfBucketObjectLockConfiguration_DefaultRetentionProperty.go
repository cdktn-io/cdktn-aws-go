package awss3


// Experimental.
type TfBucketObjectLockConfiguration_DefaultRetentionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#days TfBucketObjectLockConfiguration#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#mode TfBucketObjectLockConfiguration#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#years TfBucketObjectLockConfiguration#years}.
	// Experimental.
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

