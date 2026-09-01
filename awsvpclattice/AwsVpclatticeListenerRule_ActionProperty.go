package awsvpclattice


// Experimental.
type AwsVpclatticeListenerRule_ActionProperty struct {
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#fixed_response AwsVpclatticeListenerRule#fixed_response}
	// Experimental.
	FixedResponse *AwsVpclatticeListenerRule_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#forward AwsVpclatticeListenerRule#forward}
	// Experimental.
	Forward *AwsVpclatticeListenerRule_ForwardProperty `field:"optional" json:"forward" yaml:"forward"`
}

