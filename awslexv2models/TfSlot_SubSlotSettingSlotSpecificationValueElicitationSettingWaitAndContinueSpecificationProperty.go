package awslexv2models


// Experimental.
type TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#active TfSlot#active}.
	// Experimental.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// continue_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#continue_response TfSlot#continue_response}
	// Experimental.
	ContinueResponse interface{} `field:"optional" json:"continueResponse" yaml:"continueResponse"`
	// still_waiting_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#still_waiting_response TfSlot#still_waiting_response}
	// Experimental.
	StillWaitingResponse interface{} `field:"optional" json:"stillWaitingResponse" yaml:"stillWaitingResponse"`
	// waiting_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#waiting_response TfSlot#waiting_response}
	// Experimental.
	WaitingResponse interface{} `field:"optional" json:"waitingResponse" yaml:"waitingResponse"`
}

