package awsvpnsitetosite


// Experimental.
type DataTfGateway_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpn_gateway#name DataTfGateway#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpn_gateway#values DataTfGateway#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

