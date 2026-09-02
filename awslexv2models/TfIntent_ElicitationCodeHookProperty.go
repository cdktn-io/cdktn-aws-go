package awslexv2models


// Experimental.
type TfIntent_ElicitationCodeHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#enable_code_hook_invocation TfIntent#enable_code_hook_invocation}.
	// Experimental.
	EnableCodeHookInvocation interface{} `field:"optional" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#invocation_label TfIntent#invocation_label}.
	// Experimental.
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
}

