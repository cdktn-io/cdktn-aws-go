package awsvpnsitetosite


// Experimental.
type DataAwsVpnConnection_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpn_connection#name DataAwsVpnConnection#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpn_connection#values DataAwsVpnConnection#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

