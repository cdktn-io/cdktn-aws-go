package vpc


// Experimental.
type DataAwsPeeringConnection_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_peering_connection#name DataAwsPeeringConnection#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_peering_connection#values DataAwsPeeringConnection#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

