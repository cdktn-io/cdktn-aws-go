package awsvpclattice


// Experimental.
type AwsVpclatticeListener_DefaultActionProperty struct {
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener#fixed_response AwsVpclatticeListener#fixed_response}
	// Experimental.
	FixedResponse *AwsVpclatticeListener_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener#forward AwsVpclatticeListener#forward}
	// Experimental.
	Forward interface{} `field:"optional" json:"forward" yaml:"forward"`
}

