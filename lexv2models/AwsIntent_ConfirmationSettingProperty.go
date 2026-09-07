package lexv2models


// Experimental.
type AwsIntent_ConfirmationSettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active AwsIntent#active}.
	// Experimental.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#code_hook AwsIntent#code_hook}
	// Experimental.
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// confirmation_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#confirmation_conditional AwsIntent#confirmation_conditional}
	// Experimental.
	ConfirmationConditional interface{} `field:"optional" json:"confirmationConditional" yaml:"confirmationConditional"`
	// confirmation_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#confirmation_next_step AwsIntent#confirmation_next_step}
	// Experimental.
	ConfirmationNextStep interface{} `field:"optional" json:"confirmationNextStep" yaml:"confirmationNextStep"`
	// confirmation_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#confirmation_response AwsIntent#confirmation_response}
	// Experimental.
	ConfirmationResponse interface{} `field:"optional" json:"confirmationResponse" yaml:"confirmationResponse"`
	// declination_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#declination_conditional AwsIntent#declination_conditional}
	// Experimental.
	DeclinationConditional interface{} `field:"optional" json:"declinationConditional" yaml:"declinationConditional"`
	// declination_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#declination_next_step AwsIntent#declination_next_step}
	// Experimental.
	DeclinationNextStep interface{} `field:"optional" json:"declinationNextStep" yaml:"declinationNextStep"`
	// declination_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#declination_response AwsIntent#declination_response}
	// Experimental.
	DeclinationResponse interface{} `field:"optional" json:"declinationResponse" yaml:"declinationResponse"`
	// elicitation_code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#elicitation_code_hook AwsIntent#elicitation_code_hook}
	// Experimental.
	ElicitationCodeHook interface{} `field:"optional" json:"elicitationCodeHook" yaml:"elicitationCodeHook"`
	// failure_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#failure_conditional AwsIntent#failure_conditional}
	// Experimental.
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// failure_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#failure_next_step AwsIntent#failure_next_step}
	// Experimental.
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// failure_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#failure_response AwsIntent#failure_response}
	// Experimental.
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// prompt_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#prompt_specification AwsIntent#prompt_specification}
	// Experimental.
	PromptSpecification interface{} `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
}

