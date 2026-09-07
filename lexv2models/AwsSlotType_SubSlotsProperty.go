package lexv2models


// Experimental.
type AwsSlotType_SubSlotsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#name AwsSlotType#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot_type#slot_type_id AwsSlotType#slot_type_id}.
	// Experimental.
	SlotTypeId *string `field:"required" json:"slotTypeId" yaml:"slotTypeId"`
}

