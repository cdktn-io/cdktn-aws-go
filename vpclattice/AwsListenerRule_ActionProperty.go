package vpclattice


// Experimental.
type AwsListenerRule_ActionProperty struct {
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#fixed_response AwsListenerRule#fixed_response}
	// Experimental.
	FixedResponse *AwsListenerRule_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#forward AwsListenerRule#forward}
	// Experimental.
	Forward *AwsListenerRule_ForwardProperty `field:"optional" json:"forward" yaml:"forward"`
}

