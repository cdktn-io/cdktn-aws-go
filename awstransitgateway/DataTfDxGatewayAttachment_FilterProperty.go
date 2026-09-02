package awstransitgateway


// Experimental.
type DataTfDxGatewayAttachment_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_dx_gateway_attachment#name DataTfDxGatewayAttachment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_dx_gateway_attachment#values DataTfDxGatewayAttachment#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

