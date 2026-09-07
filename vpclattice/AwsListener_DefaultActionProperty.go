package vpclattice


// Experimental.
type AwsListener_DefaultActionProperty struct {
	// fixed_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener#fixed_response AwsListener#fixed_response}
	// Experimental.
	FixedResponse *AwsListener_FixedResponseProperty `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// forward block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener#forward AwsListener#forward}
	// Experimental.
	Forward interface{} `field:"optional" json:"forward" yaml:"forward"`
}

