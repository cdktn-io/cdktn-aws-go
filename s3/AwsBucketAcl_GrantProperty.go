package s3


// Experimental.
type AwsBucketAcl_GrantProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#permission AwsBucketAcl#permission}.
	// Experimental.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// grantee block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#grantee AwsBucketAcl#grantee}
	// Experimental.
	Grantee *AwsBucketAcl_GranteeProperty `field:"optional" json:"grantee" yaml:"grantee"`
}

