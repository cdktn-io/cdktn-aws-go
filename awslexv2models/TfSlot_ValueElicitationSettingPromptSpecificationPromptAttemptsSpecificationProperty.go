package awslexv2models


// Experimental.
type TfSlot_ValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#map_block_key TfSlot#map_block_key}.
	// Experimental.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// allowed_input_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#allowed_input_types TfSlot#allowed_input_types}
	// Experimental.
	AllowedInputTypes interface{} `field:"optional" json:"allowedInputTypes" yaml:"allowedInputTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#allow_interrupt TfSlot#allow_interrupt}.
	// Experimental.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// audio_and_dtmf_input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#audio_and_dtmf_input_specification TfSlot#audio_and_dtmf_input_specification}
	// Experimental.
	AudioAndDtmfInputSpecification interface{} `field:"optional" json:"audioAndDtmfInputSpecification" yaml:"audioAndDtmfInputSpecification"`
	// text_input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#text_input_specification TfSlot#text_input_specification}
	// Experimental.
	TextInputSpecification interface{} `field:"optional" json:"textInputSpecification" yaml:"textInputSpecification"`
}

