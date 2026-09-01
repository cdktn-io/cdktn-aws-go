package awslexv2models


// Experimental.
type AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#active AwsLexv2ModelsSlot#active}.
	// Experimental.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// continue_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#continue_response AwsLexv2ModelsSlot#continue_response}
	// Experimental.
	ContinueResponse interface{} `field:"optional" json:"continueResponse" yaml:"continueResponse"`
	// still_waiting_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#still_waiting_response AwsLexv2ModelsSlot#still_waiting_response}
	// Experimental.
	StillWaitingResponse interface{} `field:"optional" json:"stillWaitingResponse" yaml:"stillWaitingResponse"`
	// waiting_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#waiting_response AwsLexv2ModelsSlot#waiting_response}
	// Experimental.
	WaitingResponse interface{} `field:"optional" json:"waitingResponse" yaml:"waitingResponse"`
}

