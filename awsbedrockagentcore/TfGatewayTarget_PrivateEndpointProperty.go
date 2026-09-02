package awsbedrockagentcore


// Experimental.
type TfGatewayTarget_PrivateEndpointProperty struct {
	// managed_vpc_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#managed_vpc_resource TfGatewayTarget#managed_vpc_resource}
	// Experimental.
	ManagedVpcResource interface{} `field:"optional" json:"managedVpcResource" yaml:"managedVpcResource"`
	// self_managed_lattice_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#self_managed_lattice_resource TfGatewayTarget#self_managed_lattice_resource}
	// Experimental.
	SelfManagedLatticeResource interface{} `field:"optional" json:"selfManagedLatticeResource" yaml:"selfManagedLatticeResource"`
}

