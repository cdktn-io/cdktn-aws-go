package s3


// Experimental.
type AwsBucketLogging_TargetGrantProperty struct {
	// grantee block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#grantee AwsBucketLogging#grantee}
	// Experimental.
	Grantee *AwsBucketLogging_GranteeProperty `field:"required" json:"grantee" yaml:"grantee"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_logging#permission AwsBucketLogging#permission}.
	// Experimental.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

