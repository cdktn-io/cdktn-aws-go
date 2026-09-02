package awsec2imagebuilder


// Experimental.
type TfContainerRecipe_TargetRepositoryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#repository_name TfContainerRecipe#repository_name}.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#service TfContainerRecipe#service}.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
}

