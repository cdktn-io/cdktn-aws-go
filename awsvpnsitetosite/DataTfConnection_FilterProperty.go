package awsvpnsitetosite


// Experimental.
type DataTfConnection_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpn_connection#name DataTfConnection#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpn_connection#values DataTfConnection#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

