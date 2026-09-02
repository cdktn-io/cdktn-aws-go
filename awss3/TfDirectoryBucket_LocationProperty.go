package awss3


// Experimental.
type TfDirectoryBucket_LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_directory_bucket#name TfDirectoryBucket#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_directory_bucket#type TfDirectoryBucket#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

