package lexv2models


// Experimental.
type AwsIntent_InitialResponseSettingCodeHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active AwsIntent#active}.
	// Experimental.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#enable_code_hook_invocation AwsIntent#enable_code_hook_invocation}.
	// Experimental.
	EnableCodeHookInvocation interface{} `field:"required" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#invocation_label AwsIntent#invocation_label}.
	// Experimental.
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
	// post_code_hook_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#post_code_hook_specification AwsIntent#post_code_hook_specification}
	// Experimental.
	PostCodeHookSpecification interface{} `field:"optional" json:"postCodeHookSpecification" yaml:"postCodeHookSpecification"`
}

