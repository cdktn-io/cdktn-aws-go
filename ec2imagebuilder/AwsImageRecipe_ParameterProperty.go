package ec2imagebuilder


// Experimental.
type AwsImageRecipe_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#name AwsImageRecipe#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_recipe#value AwsImageRecipe#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

