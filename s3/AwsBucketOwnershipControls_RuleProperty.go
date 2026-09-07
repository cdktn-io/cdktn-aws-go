package s3


// Experimental.
type AwsBucketOwnershipControls_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_ownership_controls#object_ownership AwsBucketOwnershipControls#object_ownership}.
	// Experimental.
	ObjectOwnership *string `field:"required" json:"objectOwnership" yaml:"objectOwnership"`
}

