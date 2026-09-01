package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseMessageGroupVariationProperty struct {
	// custom_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#custom_payload AwsLexv2ModelsIntent#custom_payload}
	// Experimental.
	CustomPayload interface{} `field:"optional" json:"customPayload" yaml:"customPayload"`
	// image_response_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#image_response_card AwsLexv2ModelsIntent#image_response_card}
	// Experimental.
	ImageResponseCard interface{} `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// plain_text_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#plain_text_message AwsLexv2ModelsIntent#plain_text_message}
	// Experimental.
	PlainTextMessage interface{} `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// ssml_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#ssml_message AwsLexv2ModelsIntent#ssml_message}
	// Experimental.
	SsmlMessage interface{} `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

