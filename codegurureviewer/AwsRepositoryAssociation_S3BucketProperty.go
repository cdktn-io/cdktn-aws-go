package codegurureviewer


// Experimental.
type AwsRepositoryAssociation_S3BucketProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#bucket_name AwsRepositoryAssociation#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#name AwsRepositoryAssociation#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

