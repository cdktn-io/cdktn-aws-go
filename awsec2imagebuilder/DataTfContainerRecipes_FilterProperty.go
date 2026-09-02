package awsec2imagebuilder


// Experimental.
type DataTfContainerRecipes_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_container_recipes#name DataTfContainerRecipes#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_container_recipes#values DataTfContainerRecipes#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

