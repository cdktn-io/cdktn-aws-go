package awssagemakerai


// Experimental.
type AwsSagemakerAlgorithm_BaseModelProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#hub_content_name AwsSagemakerAlgorithm#hub_content_name}.
	// Experimental.
	HubContentName *string `field:"optional" json:"hubContentName" yaml:"hubContentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#hub_content_version AwsSagemakerAlgorithm#hub_content_version}.
	// Experimental.
	HubContentVersion *string `field:"optional" json:"hubContentVersion" yaml:"hubContentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_algorithm#recipe_name AwsSagemakerAlgorithm#recipe_name}.
	// Experimental.
	RecipeName *string `field:"optional" json:"recipeName" yaml:"recipeName"`
}

