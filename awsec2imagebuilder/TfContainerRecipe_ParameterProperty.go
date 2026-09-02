package awsec2imagebuilder


// Experimental.
type TfContainerRecipe_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#name TfContainerRecipe#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_container_recipe#value TfContainerRecipe#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

