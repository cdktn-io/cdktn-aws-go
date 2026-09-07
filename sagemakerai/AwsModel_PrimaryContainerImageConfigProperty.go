package sagemakerai


// Experimental.
type AwsModel_PrimaryContainerImageConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#repository_access_mode AwsModel#repository_access_mode}.
	// Experimental.
	RepositoryAccessMode *string `field:"required" json:"repositoryAccessMode" yaml:"repositoryAccessMode"`
	// repository_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_model#repository_auth_config AwsModel#repository_auth_config}
	// Experimental.
	RepositoryAuthConfig *AwsModel_PrimaryContainerImageConfigRepositoryAuthConfigProperty `field:"optional" json:"repositoryAuthConfig" yaml:"repositoryAuthConfig"`
}

