package lexmodelbuilding


// Experimental.
type AwsIntent_ConclusionStatementMessageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#content AwsIntent#content}.
	// Experimental.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#content_type AwsIntent#content_type}.
	// Experimental.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#group_number AwsIntent#group_number}.
	// Experimental.
	GroupNumber *float64 `field:"optional" json:"groupNumber" yaml:"groupNumber"`
}

