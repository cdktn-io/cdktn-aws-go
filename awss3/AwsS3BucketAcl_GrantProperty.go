package awss3


// Experimental.
type AwsS3BucketAcl_GrantProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#permission AwsS3BucketAcl#permission}.
	// Experimental.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// grantee block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#grantee AwsS3BucketAcl#grantee}
	// Experimental.
	Grantee *AwsS3BucketAcl_GranteeProperty `field:"optional" json:"grantee" yaml:"grantee"`
}

