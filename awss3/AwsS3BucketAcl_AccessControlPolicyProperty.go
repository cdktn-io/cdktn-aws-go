package awss3


// Experimental.
type AwsS3BucketAcl_AccessControlPolicyProperty struct {
	// owner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#owner AwsS3BucketAcl#owner}
	// Experimental.
	Owner *AwsS3BucketAcl_OwnerProperty `field:"required" json:"owner" yaml:"owner"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_acl#grant AwsS3BucketAcl#grant}
	// Experimental.
	Grant interface{} `field:"optional" json:"grant" yaml:"grant"`
}

