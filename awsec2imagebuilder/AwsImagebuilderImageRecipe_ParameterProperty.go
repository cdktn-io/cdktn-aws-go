package awsec2imagebuilder


// Experimental.
type AwsImagebuilderImageRecipe_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#name AwsImagebuilderImageRecipe#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#value AwsImagebuilderImageRecipe#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

