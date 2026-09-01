package awstransitgateway


// Experimental.
type DataAwsEc2TransitGatewayRouteTable_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_route_table#name DataAwsEc2TransitGatewayRouteTable#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_route_table#values DataAwsEc2TransitGatewayRouteTable#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

