package awssavingsplans


// Experimental.
type DataAwsSavingsplansOfferings_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#name DataAwsSavingsplansOfferings#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#values DataAwsSavingsplansOfferings#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

