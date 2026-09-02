package awslexmodelbuilding


// Experimental.
type TfBot_AbortStatementProperty struct {
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#message TfBot#message}
	// Experimental.
	Message interface{} `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#response_card TfBot#response_card}.
	// Experimental.
	ResponseCard *string `field:"optional" json:"responseCard" yaml:"responseCard"`
}

