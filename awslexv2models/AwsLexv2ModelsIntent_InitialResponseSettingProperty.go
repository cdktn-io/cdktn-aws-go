package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_InitialResponseSettingProperty struct {
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#code_hook AwsLexv2ModelsIntent#code_hook}
	// Experimental.
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#conditional AwsLexv2ModelsIntent#conditional}
	// Experimental.
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// initial_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#initial_response AwsLexv2ModelsIntent#initial_response}
	// Experimental.
	InitialResponse interface{} `field:"optional" json:"initialResponse" yaml:"initialResponse"`
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#next_step AwsLexv2ModelsIntent#next_step}
	// Experimental.
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
}

