package awslexv2models


// Experimental.
type TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationDtmfSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#deletion_character TfSlot#deletion_character}.
	// Experimental.
	DeletionCharacter *string `field:"required" json:"deletionCharacter" yaml:"deletionCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#end_character TfSlot#end_character}.
	// Experimental.
	EndCharacter *string `field:"required" json:"endCharacter" yaml:"endCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#end_timeout_ms TfSlot#end_timeout_ms}.
	// Experimental.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#max_length TfSlot#max_length}.
	// Experimental.
	MaxLength *float64 `field:"required" json:"maxLength" yaml:"maxLength"`
}

