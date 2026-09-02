package awslexv2models


// Experimental.
type TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingProperty struct {
	// default_value_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#default_value_specification TfSlot#default_value_specification}
	// Experimental.
	DefaultValueSpecification interface{} `field:"optional" json:"defaultValueSpecification" yaml:"defaultValueSpecification"`
	// prompt_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#prompt_specification TfSlot#prompt_specification}
	// Experimental.
	PromptSpecification interface{} `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
	// sample_utterance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#sample_utterance TfSlot#sample_utterance}
	// Experimental.
	SampleUtterance interface{} `field:"optional" json:"sampleUtterance" yaml:"sampleUtterance"`
	// wait_and_continue_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#wait_and_continue_specification TfSlot#wait_and_continue_specification}
	// Experimental.
	WaitAndContinueSpecification interface{} `field:"optional" json:"waitAndContinueSpecification" yaml:"waitAndContinueSpecification"`
}

