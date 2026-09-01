package awss3


// Experimental.
type AwsS3BucketLogging_TargetGrantProperty struct {
	// grantee block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#grantee AwsS3BucketLogging#grantee}
	// Experimental.
	Grantee *AwsS3BucketLogging_GranteeProperty `field:"required" json:"grantee" yaml:"grantee"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#permission AwsS3BucketLogging#permission}.
	// Experimental.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

