package awsvpclattice


// Experimental.
type TfListenerRule_ActionProperty struct {
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#fixed_response TfListenerRule#fixed_response}
	// Experimental.
	FixedResponse *TfListenerRule_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#forward TfListenerRule#forward}
	// Experimental.
	Forward *TfListenerRule_ForwardProperty `field:"optional" json:"forward" yaml:"forward"`
}

