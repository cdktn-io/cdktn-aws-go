package awslexv2models


// Experimental.
type TfIntent_ConfirmationSettingDeclinationConditionalDefaultBranchNextStepDialogActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#type TfIntent#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#slot_to_elicit TfIntent#slot_to_elicit}.
	// Experimental.
	SlotToElicit *string `field:"optional" json:"slotToElicit" yaml:"slotToElicit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#suppress_next_message TfIntent#suppress_next_message}.
	// Experimental.
	SuppressNextMessage interface{} `field:"optional" json:"suppressNextMessage" yaml:"suppressNextMessage"`
}

