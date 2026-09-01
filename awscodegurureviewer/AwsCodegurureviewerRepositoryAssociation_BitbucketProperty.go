package awscodegurureviewer


// Experimental.
type AwsCodegurureviewerRepositoryAssociation_BitbucketProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#connection_arn AwsCodegurureviewerRepositoryAssociation#connection_arn}.
	// Experimental.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#name AwsCodegurureviewerRepositoryAssociation#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#owner AwsCodegurureviewerRepositoryAssociation#owner}.
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
}

