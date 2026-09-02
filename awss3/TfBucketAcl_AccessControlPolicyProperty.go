package awss3


// Experimental.
type TfBucketAcl_AccessControlPolicyProperty struct {
	// owner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#owner TfBucketAcl#owner}
	// Experimental.
	Owner *TfBucketAcl_OwnerProperty `field:"required" json:"owner" yaml:"owner"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#grant TfBucketAcl#grant}
	// Experimental.
	Grant interface{} `field:"optional" json:"grant" yaml:"grant"`
}

