package s3


// Experimental.
type AwsDirectoryBucket_LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_directory_bucket#name AwsDirectoryBucket#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_directory_bucket#type AwsDirectoryBucket#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

