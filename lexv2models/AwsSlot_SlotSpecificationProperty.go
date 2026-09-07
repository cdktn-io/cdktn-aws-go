package lexv2models


// Experimental.
type AwsSlot_SlotSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#map_block_key AwsSlot#map_block_key}.
	// Experimental.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#slot_type_id AwsSlot#slot_type_id}.
	// Experimental.
	SlotTypeId *string `field:"required" json:"slotTypeId" yaml:"slotTypeId"`
	// value_elicitation_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#value_elicitation_setting AwsSlot#value_elicitation_setting}
	// Experimental.
	ValueElicitationSetting interface{} `field:"optional" json:"valueElicitationSetting" yaml:"valueElicitationSetting"`
}

