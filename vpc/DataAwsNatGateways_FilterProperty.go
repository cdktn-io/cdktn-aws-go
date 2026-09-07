package vpc


// Experimental.
type DataAwsNatGateways_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/nat_gateways#name DataAwsNatGateways#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/nat_gateways#values DataAwsNatGateways#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

