package awsnetworkmanager


// Experimental.
type AwsNetworkmanagerTransitGatewayRouteTableAttachment_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#create AwsNetworkmanagerTransitGatewayRouteTableAttachment#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#delete AwsNetworkmanagerTransitGatewayRouteTableAttachment#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

