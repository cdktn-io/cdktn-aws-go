package lexv2models


// Experimental.
type AwsSlot_ValueElicitationSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#slot_constraint AwsSlot#slot_constraint}.
	// Experimental.
	SlotConstraint *string `field:"required" json:"slotConstraint" yaml:"slotConstraint"`
	// default_value_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#default_value_specification AwsSlot#default_value_specification}
	// Experimental.
	DefaultValueSpecification interface{} `field:"optional" json:"defaultValueSpecification" yaml:"defaultValueSpecification"`
	// prompt_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#prompt_specification AwsSlot#prompt_specification}
	// Experimental.
	PromptSpecification interface{} `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
	// sample_utterance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#sample_utterance AwsSlot#sample_utterance}
	// Experimental.
	SampleUtterance interface{} `field:"optional" json:"sampleUtterance" yaml:"sampleUtterance"`
	// slot_resolution_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#slot_resolution_setting AwsSlot#slot_resolution_setting}
	// Experimental.
	SlotResolutionSetting interface{} `field:"optional" json:"slotResolutionSetting" yaml:"slotResolutionSetting"`
	// wait_and_continue_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#wait_and_continue_specification AwsSlot#wait_and_continue_specification}
	// Experimental.
	WaitAndContinueSpecification interface{} `field:"optional" json:"waitAndContinueSpecification" yaml:"waitAndContinueSpecification"`
}

