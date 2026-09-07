package ec2imagebuilder


// Experimental.
type DataAwsImageRecipes_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_image_recipes#name DataAwsImageRecipes#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_image_recipes#values DataAwsImageRecipes#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

