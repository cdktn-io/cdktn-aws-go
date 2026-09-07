package vpnsitetosite


// Experimental.
type DataAwsConnection_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpn_connection#name DataAwsConnection#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpn_connection#values DataAwsConnection#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

