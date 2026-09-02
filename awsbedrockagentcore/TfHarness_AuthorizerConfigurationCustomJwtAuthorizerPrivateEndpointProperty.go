package awsbedrockagentcore


// Experimental.
type TfHarness_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointProperty struct {
	// managed_vpc_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#managed_vpc_resource TfHarness#managed_vpc_resource}
	// Experimental.
	ManagedVpcResource interface{} `field:"optional" json:"managedVpcResource" yaml:"managedVpcResource"`
	// self_managed_lattice_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#self_managed_lattice_resource TfHarness#self_managed_lattice_resource}
	// Experimental.
	SelfManagedLatticeResource interface{} `field:"optional" json:"selfManagedLatticeResource" yaml:"selfManagedLatticeResource"`
}

