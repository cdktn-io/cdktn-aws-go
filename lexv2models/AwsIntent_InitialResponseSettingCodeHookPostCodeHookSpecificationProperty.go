package lexv2models


// Experimental.
type AwsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationProperty struct {
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
	// success_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#success_conditional AwsIntent#success_conditional}
	// Experimental.
	SuccessConditional interface{} `field:"optional" json:"successConditional" yaml:"successConditional"`
	// success_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#success_next_step AwsIntent#success_next_step}
	// Experimental.
	SuccessNextStep interface{} `field:"optional" json:"successNextStep" yaml:"successNextStep"`
	// success_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#success_response AwsIntent#success_response}
	// Experimental.
	SuccessResponse interface{} `field:"optional" json:"successResponse" yaml:"successResponse"`
	// timeout_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#timeout_conditional AwsIntent#timeout_conditional}
	// Experimental.
	TimeoutConditional interface{} `field:"optional" json:"timeoutConditional" yaml:"timeoutConditional"`
	// timeout_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#timeout_next_step AwsIntent#timeout_next_step}
	// Experimental.
	TimeoutNextStep interface{} `field:"optional" json:"timeoutNextStep" yaml:"timeoutNextStep"`
	// timeout_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#timeout_response AwsIntent#timeout_response}
	// Experimental.
	TimeoutResponse interface{} `field:"optional" json:"timeoutResponse" yaml:"timeoutResponse"`
}

