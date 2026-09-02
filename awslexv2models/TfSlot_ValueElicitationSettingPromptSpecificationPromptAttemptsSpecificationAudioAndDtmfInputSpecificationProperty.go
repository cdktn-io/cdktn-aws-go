package awslexv2models


// Experimental.
type TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#start_timeout_ms TfSlot#start_timeout_ms}.
	// Experimental.
	StartTimeoutMs *float64 `field:"required" json:"startTimeoutMs" yaml:"startTimeoutMs"`
	// audio_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#audio_specification TfSlot#audio_specification}
	// Experimental.
	AudioSpecification interface{} `field:"optional" json:"audioSpecification" yaml:"audioSpecification"`
	// dtmf_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#dtmf_specification TfSlot#dtmf_specification}
	// Experimental.
	DtmfSpecification interface{} `field:"optional" json:"dtmfSpecification" yaml:"dtmfSpecification"`
}

