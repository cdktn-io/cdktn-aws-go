package transitgateway


// Experimental.
type DataAwsPeeringAttachment_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_peering_attachment#name DataAwsPeeringAttachment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_peering_attachment#values DataAwsPeeringAttachment#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

