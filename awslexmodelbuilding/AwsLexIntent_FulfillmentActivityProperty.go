package awslexmodelbuilding


// Experimental.
type AwsLexIntent_FulfillmentActivityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#type AwsLexIntent#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#code_hook AwsLexIntent#code_hook}
	// Experimental.
	CodeHook *AwsLexIntent_CodeHookProperty `field:"optional" json:"codeHook" yaml:"codeHook"`
}

