package awslexmodelbuilding


// Experimental.
type TfIntent_ConfirmationPromptMessageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#content TfIntent#content}.
	// Experimental.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#content_type TfIntent#content_type}.
	// Experimental.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#group_number TfIntent#group_number}.
	// Experimental.
	GroupNumber *float64 `field:"optional" json:"groupNumber" yaml:"groupNumber"`
}

