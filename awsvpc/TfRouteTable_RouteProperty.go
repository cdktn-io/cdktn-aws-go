package awsvpc


// Experimental.
type TfRouteTable_RouteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#carrier_gateway_id TfRouteTable#carrier_gateway_id}.
	// Experimental.
	CarrierGatewayId *string `field:"optional" json:"carrierGatewayId" yaml:"carrierGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#cidr_block TfRouteTable#cidr_block}.
	// Experimental.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#core_network_arn TfRouteTable#core_network_arn}.
	// Experimental.
	CoreNetworkArn *string `field:"optional" json:"coreNetworkArn" yaml:"coreNetworkArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#destination_prefix_list_id TfRouteTable#destination_prefix_list_id}.
	// Experimental.
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#egress_only_gateway_id TfRouteTable#egress_only_gateway_id}.
	// Experimental.
	EgressOnlyGatewayId *string `field:"optional" json:"egressOnlyGatewayId" yaml:"egressOnlyGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#gateway_id TfRouteTable#gateway_id}.
	// Experimental.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#ipv6_cidr_block TfRouteTable#ipv6_cidr_block}.
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#local_gateway_id TfRouteTable#local_gateway_id}.
	// Experimental.
	LocalGatewayId *string `field:"optional" json:"localGatewayId" yaml:"localGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#nat_gateway_id TfRouteTable#nat_gateway_id}.
	// Experimental.
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#network_interface_id TfRouteTable#network_interface_id}.
	// Experimental.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#odb_network_arn TfRouteTable#odb_network_arn}.
	// Experimental.
	OdbNetworkArn *string `field:"optional" json:"odbNetworkArn" yaml:"odbNetworkArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#transit_gateway_id TfRouteTable#transit_gateway_id}.
	// Experimental.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#vpc_endpoint_id TfRouteTable#vpc_endpoint_id}.
	// Experimental.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route_table#vpc_peering_connection_id TfRouteTable#vpc_peering_connection_id}.
	// Experimental.
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

