package awssavingsplans


// Experimental.
type DataTfOfferings_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#name DataTfOfferings#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#values DataTfOfferings#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

