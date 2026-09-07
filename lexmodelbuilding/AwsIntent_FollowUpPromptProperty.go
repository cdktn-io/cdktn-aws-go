package lexmodelbuilding


// Experimental.
type AwsIntent_FollowUpPromptProperty struct {
	// prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#prompt AwsIntent#prompt}
	// Experimental.
	Prompt *AwsIntent_PromptProperty `field:"required" json:"prompt" yaml:"prompt"`
	// rejection_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#rejection_statement AwsIntent#rejection_statement}
	// Experimental.
	RejectionStatement *AwsIntent_FollowUpPromptRejectionStatementProperty `field:"required" json:"rejectionStatement" yaml:"rejectionStatement"`
}

