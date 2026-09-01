package awssagemakerai


// Experimental.
type AwsSagemakerModel_ContainerImageConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#repository_access_mode AwsSagemakerModel#repository_access_mode}.
	// Experimental.
	RepositoryAccessMode *string `field:"required" json:"repositoryAccessMode" yaml:"repositoryAccessMode"`
	// repository_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#repository_auth_config AwsSagemakerModel#repository_auth_config}
	// Experimental.
	RepositoryAuthConfig *AwsSagemakerModel_ContainerImageConfigRepositoryAuthConfigProperty `field:"optional" json:"repositoryAuthConfig" yaml:"repositoryAuthConfig"`
}

