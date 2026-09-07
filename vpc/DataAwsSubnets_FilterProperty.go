package vpc


// Experimental.
type DataAwsSubnets_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/subnets#name DataAwsSubnets#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/subnets#values DataAwsSubnets#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

