package s3


// Experimental.
type AwsBucket_ReplicationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#role AwsBucket#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#rules AwsBucket#rules}
	// Experimental.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

