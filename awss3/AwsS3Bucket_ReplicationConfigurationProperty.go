package awss3


// Experimental.
type AwsS3Bucket_ReplicationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#role AwsS3Bucket#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#rules AwsS3Bucket#rules}
	// Experimental.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

