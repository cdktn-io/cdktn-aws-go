package awsvpc


// Experimental.
type DataTfNatGateways_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/nat_gateways#name DataTfNatGateways#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/nat_gateways#values DataTfNatGateways#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

