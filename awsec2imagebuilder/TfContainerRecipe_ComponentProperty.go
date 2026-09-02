package awsec2imagebuilder


// Experimental.
type TfContainerRecipe_ComponentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#component_arn TfContainerRecipe#component_arn}.
	// Experimental.
	ComponentArn *string `field:"required" json:"componentArn" yaml:"componentArn"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#parameter TfContainerRecipe#parameter}
	// Experimental.
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
}

