package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_InitialResponseSettingConditionalDefaultBranchNextStepDialogActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#type AwsLexv2ModelsIntent#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#slot_to_elicit AwsLexv2ModelsIntent#slot_to_elicit}.
	// Experimental.
	SlotToElicit *string `field:"optional" json:"slotToElicit" yaml:"slotToElicit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#suppress_next_message AwsLexv2ModelsIntent#suppress_next_message}.
	// Experimental.
	SuppressNextMessage interface{} `field:"optional" json:"suppressNextMessage" yaml:"suppressNextMessage"`
}

