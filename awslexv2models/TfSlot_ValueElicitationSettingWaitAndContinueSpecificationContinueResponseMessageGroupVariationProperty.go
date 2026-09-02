package awslexv2models


// Experimental.
type TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty struct {
	// custom_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#custom_payload TfSlot#custom_payload}
	// Experimental.
	CustomPayload interface{} `field:"optional" json:"customPayload" yaml:"customPayload"`
	// image_response_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#image_response_card TfSlot#image_response_card}
	// Experimental.
	ImageResponseCard interface{} `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// plain_text_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#plain_text_message TfSlot#plain_text_message}
	// Experimental.
	PlainTextMessage interface{} `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// ssml_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#ssml_message TfSlot#ssml_message}
	// Experimental.
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

