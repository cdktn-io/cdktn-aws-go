package awscodegurureviewer


// Experimental.
type TfRepositoryAssociation_S3BucketProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#bucket_name TfRepositoryAssociation#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#name TfRepositoryAssociation#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

