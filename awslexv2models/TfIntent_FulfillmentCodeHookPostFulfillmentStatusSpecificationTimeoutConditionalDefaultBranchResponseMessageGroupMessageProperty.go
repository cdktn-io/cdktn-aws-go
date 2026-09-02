package awslexv2models


// Experimental.
type TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupMessageProperty struct {
	// custom_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#custom_payload TfIntent#custom_payload}
	// Experimental.
	CustomPayload interface{} `field:"optional" json:"customPayload" yaml:"customPayload"`
	// image_response_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#image_response_card TfIntent#image_response_card}
	// Experimental.
	ImageResponseCard interface{} `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// plain_text_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#plain_text_message TfIntent#plain_text_message}
	// Experimental.
	PlainTextMessage interface{} `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// ssml_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#ssml_message TfIntent#ssml_message}
	// Experimental.
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

