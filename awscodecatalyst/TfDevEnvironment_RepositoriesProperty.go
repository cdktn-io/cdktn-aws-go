package awscodecatalyst


// Experimental.
type TfDevEnvironment_RepositoriesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codecatalyst_dev_environment#repository_name TfDevEnvironment#repository_name}.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codecatalyst_dev_environment#branch_name TfDevEnvironment#branch_name}.
	// Experimental.
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
}

