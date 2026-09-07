package ec2imagebuilder


// Experimental.
type DataAwsContainerRecipes_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_container_recipes#name DataAwsContainerRecipes#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_container_recipes#values DataAwsContainerRecipes#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

