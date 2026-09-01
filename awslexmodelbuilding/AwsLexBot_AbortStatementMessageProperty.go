package awslexmodelbuilding


// Experimental.
type AwsLexBot_AbortStatementMessageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#content AwsLexBot#content}.
	// Experimental.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#content_type AwsLexBot#content_type}.
	// Experimental.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#group_number AwsLexBot#group_number}.
	// Experimental.
	GroupNumber *float64 `field:"optional" json:"groupNumber" yaml:"groupNumber"`
}

