package awslexmodelbuilding


// Experimental.
type TfIntent_FollowUpPromptProperty struct {
	// prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#prompt TfIntent#prompt}
	// Experimental.
	Prompt *TfIntent_PromptProperty `field:"required" json:"prompt" yaml:"prompt"`
	// rejection_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#rejection_statement TfIntent#rejection_statement}
	// Experimental.
	RejectionStatement *TfIntent_FollowUpPromptRejectionStatementProperty `field:"required" json:"rejectionStatement" yaml:"rejectionStatement"`
}

