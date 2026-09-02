package awslexv2models


// Experimental.
type TfSlot_SubSlotSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#expression TfSlot#expression}.
	// Experimental.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// slot_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#slot_specification TfSlot#slot_specification}
	// Experimental.
	SlotSpecification interface{} `field:"optional" json:"slotSpecification" yaml:"slotSpecification"`
}

