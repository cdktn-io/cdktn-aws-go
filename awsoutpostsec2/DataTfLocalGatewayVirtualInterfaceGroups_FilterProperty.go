package awsoutpostsec2


// Experimental.
type DataTfLocalGatewayVirtualInterfaceGroups_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_local_gateway_virtual_interface_groups#name DataTfLocalGatewayVirtualInterfaceGroups#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_local_gateway_virtual_interface_groups#values DataTfLocalGatewayVirtualInterfaceGroups#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

