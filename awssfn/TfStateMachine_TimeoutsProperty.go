package awssfn


// Experimental.
type TfStateMachine_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_state_machine#create TfStateMachine#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_state_machine#delete TfStateMachine#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_state_machine#update TfStateMachine#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

