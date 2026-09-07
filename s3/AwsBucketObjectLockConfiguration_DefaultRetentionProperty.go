package s3


// Experimental.
type AwsBucketObjectLockConfiguration_DefaultRetentionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#days AwsBucketObjectLockConfiguration#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#mode AwsBucketObjectLockConfiguration#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_object_lock_configuration#years AwsBucketObjectLockConfiguration#years}.
	// Experimental.
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

