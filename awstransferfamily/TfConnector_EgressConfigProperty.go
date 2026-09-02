package awstransferfamily


// Experimental.
type TfConnector_EgressConfigProperty struct {
	// vpc_lattice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#vpc_lattice TfConnector#vpc_lattice}
	// Experimental.
	VpcLattice *TfConnector_VpcLatticeProperty `field:"optional" json:"vpcLattice" yaml:"vpcLattice"`
}

