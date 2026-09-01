package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_ClosingSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active AwsLexv2ModelsIntent#active}.
	// Experimental.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// closing_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#closing_response AwsLexv2ModelsIntent#closing_response}
	// Experimental.
	ClosingResponse interface{} `field:"optional" json:"closingResponse" yaml:"closingResponse"`
	// conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#conditional AwsLexv2ModelsIntent#conditional}
	// Experimental.
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#next_step AwsLexv2ModelsIntent#next_step}
	// Experimental.
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
}

