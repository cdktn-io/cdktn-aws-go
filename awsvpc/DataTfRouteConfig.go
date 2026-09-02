package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfRouteConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#route_table_id DataTfRoute#route_table_id}.
	// Experimental.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#carrier_gateway_id DataTfRoute#carrier_gateway_id}.
	// Experimental.
	CarrierGatewayId *string `field:"optional" json:"carrierGatewayId" yaml:"carrierGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#core_network_arn DataTfRoute#core_network_arn}.
	// Experimental.
	CoreNetworkArn *string `field:"optional" json:"coreNetworkArn" yaml:"coreNetworkArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#destination_cidr_block DataTfRoute#destination_cidr_block}.
	// Experimental.
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#destination_ipv6_cidr_block DataTfRoute#destination_ipv6_cidr_block}.
	// Experimental.
	DestinationIpv6CidrBlock *string `field:"optional" json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#destination_prefix_list_id DataTfRoute#destination_prefix_list_id}.
	// Experimental.
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#egress_only_gateway_id DataTfRoute#egress_only_gateway_id}.
	// Experimental.
	EgressOnlyGatewayId *string `field:"optional" json:"egressOnlyGatewayId" yaml:"egressOnlyGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#gateway_id DataTfRoute#gateway_id}.
	// Experimental.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#id DataTfRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#instance_id DataTfRoute#instance_id}.
	// Experimental.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#local_gateway_id DataTfRoute#local_gateway_id}.
	// Experimental.
	LocalGatewayId *string `field:"optional" json:"localGatewayId" yaml:"localGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#nat_gateway_id DataTfRoute#nat_gateway_id}.
	// Experimental.
	NatGatewayId *string `field:"optional" json:"natGatewayId" yaml:"natGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#network_interface_id DataTfRoute#network_interface_id}.
	// Experimental.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#odb_network_arn DataTfRoute#odb_network_arn}.
	// Experimental.
	OdbNetworkArn *string `field:"optional" json:"odbNetworkArn" yaml:"odbNetworkArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#region DataTfRoute#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#timeouts DataTfRoute#timeouts}
	// Experimental.
	Timeouts *DataTfRoute_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#transit_gateway_id DataTfRoute#transit_gateway_id}.
	// Experimental.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route#vpc_peering_connection_id DataTfRoute#vpc_peering_connection_id}.
	// Experimental.
	VpcPeeringConnectionId *string `field:"optional" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
}

