package awsec2imagebuilder


// Experimental.
type DataTfComponents_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_components#name DataTfComponents#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/imagebuilder_components#values DataTfComponents#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

