package s3


// Experimental.
type AwsBucketAcl_AccessControlPolicyProperty struct {
	// owner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#owner AwsBucketAcl#owner}
	// Experimental.
	Owner *AwsBucketAcl_OwnerProperty `field:"required" json:"owner" yaml:"owner"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#grant AwsBucketAcl#grant}
	// Experimental.
	Grant interface{} `field:"optional" json:"grant" yaml:"grant"`
}

