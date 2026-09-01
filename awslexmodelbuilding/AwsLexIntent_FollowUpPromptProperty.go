package awslexmodelbuilding


// Experimental.
type AwsLexIntent_FollowUpPromptProperty struct {
	// prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#prompt AwsLexIntent#prompt}
	// Experimental.
	Prompt *AwsLexIntent_PromptProperty `field:"required" json:"prompt" yaml:"prompt"`
	// rejection_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#rejection_statement AwsLexIntent#rejection_statement}
	// Experimental.
	RejectionStatement *AwsLexIntent_FollowUpPromptRejectionStatementProperty `field:"required" json:"rejectionStatement" yaml:"rejectionStatement"`
}

