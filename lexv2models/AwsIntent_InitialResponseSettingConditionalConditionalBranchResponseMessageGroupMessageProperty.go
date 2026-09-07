package lexv2models


// Experimental.
type AwsIntent_InitialResponseSettingConditionalConditionalBranchResponseMessageGroupMessageProperty struct {
	// custom_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#custom_payload AwsIntent#custom_payload}
	// Experimental.
	CustomPayload interface{} `field:"optional" json:"customPayload" yaml:"customPayload"`
	// image_response_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#image_response_card AwsIntent#image_response_card}
	// Experimental.
	ImageResponseCard interface{} `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// plain_text_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#plain_text_message AwsIntent#plain_text_message}
	// Experimental.
	PlainTextMessage interface{} `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// ssml_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#ssml_message AwsIntent#ssml_message}
	// Experimental.
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

