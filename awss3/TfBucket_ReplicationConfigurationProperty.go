package awss3


// Experimental.
type TfBucket_ReplicationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#role TfBucket#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#rules TfBucket#rules}
	// Experimental.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

