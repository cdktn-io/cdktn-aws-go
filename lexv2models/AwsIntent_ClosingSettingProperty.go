package lexv2models


// Experimental.
type AwsIntent_ClosingSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active AwsIntent#active}.
	// Experimental.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// closing_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#closing_response AwsIntent#closing_response}
	// Experimental.
	ClosingResponse interface{} `field:"optional" json:"closingResponse" yaml:"closingResponse"`
	// conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#conditional AwsIntent#conditional}
	// Experimental.
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#next_step AwsIntent#next_step}
	// Experimental.
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
}

