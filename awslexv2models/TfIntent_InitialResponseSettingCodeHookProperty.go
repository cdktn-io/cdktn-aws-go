package awslexv2models


// Experimental.
type TfIntent_InitialResponseSettingCodeHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active TfIntent#active}.
	// Experimental.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#enable_code_hook_invocation TfIntent#enable_code_hook_invocation}.
	// Experimental.
	EnableCodeHookInvocation interface{} `field:"required" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#invocation_label TfIntent#invocation_label}.
	// Experimental.
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
	// post_code_hook_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#post_code_hook_specification TfIntent#post_code_hook_specification}
	// Experimental.
	PostCodeHookSpecification interface{} `field:"optional" json:"postCodeHookSpecification" yaml:"postCodeHookSpecification"`
}

