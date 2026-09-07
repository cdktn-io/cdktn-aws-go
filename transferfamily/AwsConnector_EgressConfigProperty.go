package transferfamily


// Experimental.
type AwsConnector_EgressConfigProperty struct {
	// vpc_lattice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#vpc_lattice AwsConnector#vpc_lattice}
	// Experimental.
	VpcLattice *AwsConnector_VpcLatticeProperty `field:"optional" json:"vpcLattice" yaml:"vpcLattice"`
}

