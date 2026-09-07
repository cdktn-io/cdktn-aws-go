package outpostsec2


// Experimental.
type DataAwsLocalGatewayVirtualInterface_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_local_gateway_virtual_interface#name DataAwsLocalGatewayVirtualInterface#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_local_gateway_virtual_interface#values DataAwsLocalGatewayVirtualInterface#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

