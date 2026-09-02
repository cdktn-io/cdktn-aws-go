package awsvpc


// Experimental.
type TfDefaultRouteTable_RouteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#cidr_block TfDefaultRouteTable#cidr_block}.
	// Experimental.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#core_network_arn TfDefaultRouteTable#core_network_arn}.
	// Experimental.
	CoreNetworkArn *string `field:"optional" json:"coreNetworkArn" yaml:"coreNetworkArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#destination_prefix_list_id TfDefaultRouteTable#destination_prefix_list_id}.
	// Experimental.
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#egress_only_gateway_id TfDefaultRouteTable#egress_only_gateway_id}.
	// Experimental.
	EgressOnlyGatewayId *string `field:"optional" json:"egressOnlyGatewayId" yaml:"egressOnlyGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#gateway_id TfDefaultRouteTable#gateway_id}.
	// Experimental.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#instance_id TfDefaultRouteTable#instance_id}.
	// Experimental.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#ipv6_cidr_block TfDefaultRouteTable#ipv6_cidr_block}.
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#nat_gateway_id TfDefaultRouteTable#nat_gateway_id}.
	// Experimental.
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#network_interface_id TfDefaultRouteTable#network_interface_id}.
	// Experimental.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#transit_gateway_id TfDefaultRouteTable#transit_gateway_id}.
	// Experimental.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#vpc_endpoint_id TfDefaultRouteTable#vpc_endpoint_id}.
	// Experimental.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_route_table#vpc_peering_connection_id TfDefaultRouteTable#vpc_peering_connection_id}.
	// Experimental.
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

