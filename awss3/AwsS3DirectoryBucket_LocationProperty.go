package awss3


// Experimental.
type AwsS3DirectoryBucket_LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_directory_bucket#name AwsS3DirectoryBucket#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_directory_bucket#type AwsS3DirectoryBucket#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

