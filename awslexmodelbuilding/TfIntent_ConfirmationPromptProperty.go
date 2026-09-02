package awslexmodelbuilding


// Experimental.
type TfIntent_ConfirmationPromptProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#max_attempts TfIntent#max_attempts}.
	// Experimental.
	MaxAttempts *float64 `field:"required" json:"maxAttempts" yaml:"maxAttempts"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#message TfIntent#message}
	// Experimental.
	Message interface{} `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#response_card TfIntent#response_card}.
	// Experimental.
	ResponseCard *string `field:"optional" json:"responseCard" yaml:"responseCard"`
}

