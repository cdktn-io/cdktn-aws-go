package awslexv2models


// Experimental.
type TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationAudioSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#end_timeout_ms TfSlot#end_timeout_ms}.
	// Experimental.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#max_length_ms TfSlot#max_length_ms}.
	// Experimental.
	MaxLengthMs *float64 `field:"required" json:"maxLengthMs" yaml:"maxLengthMs"`
}

