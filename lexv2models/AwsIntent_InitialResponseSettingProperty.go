package lexv2models


// Experimental.
type AwsIntent_InitialResponseSettingProperty struct {
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#code_hook AwsIntent#code_hook}
	// Experimental.
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#conditional AwsIntent#conditional}
	// Experimental.
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// initial_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#initial_response AwsIntent#initial_response}
	// Experimental.
	InitialResponse interface{} `field:"optional" json:"initialResponse" yaml:"initialResponse"`
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#next_step AwsIntent#next_step}
	// Experimental.
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
}

