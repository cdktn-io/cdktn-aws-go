package lexmodelbuilding


// Experimental.
type AwsIntent_FulfillmentActivityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#type AwsIntent#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#code_hook AwsIntent#code_hook}
	// Experimental.
	CodeHook *AwsIntent_CodeHookProperty `field:"optional" json:"codeHook" yaml:"codeHook"`
}

