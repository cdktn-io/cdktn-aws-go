package awslexv2models


// Experimental.
type TfIntent_ClosingSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active TfIntent#active}.
	// Experimental.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// closing_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#closing_response TfIntent#closing_response}
	// Experimental.
	ClosingResponse interface{} `field:"optional" json:"closingResponse" yaml:"closingResponse"`
	// conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#conditional TfIntent#conditional}
	// Experimental.
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#next_step TfIntent#next_step}
	// Experimental.
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
}

