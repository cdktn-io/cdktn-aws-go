package awstransitgateway


// Experimental.
type DataTfPeeringAttachment_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_peering_attachment#name DataTfPeeringAttachment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_peering_attachment#values DataTfPeeringAttachment#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

