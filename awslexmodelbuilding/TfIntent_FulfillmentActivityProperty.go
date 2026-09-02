package awslexmodelbuilding


// Experimental.
type TfIntent_FulfillmentActivityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#type TfIntent#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#code_hook TfIntent#code_hook}
	// Experimental.
	CodeHook *TfIntent_CodeHookProperty `field:"optional" json:"codeHook" yaml:"codeHook"`
}

