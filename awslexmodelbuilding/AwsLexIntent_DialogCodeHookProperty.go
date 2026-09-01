package awslexmodelbuilding


// Experimental.
type AwsLexIntent_DialogCodeHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#message_version AwsLexIntent#message_version}.
	// Experimental.
	MessageVersion *string `field:"required" json:"messageVersion" yaml:"messageVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#uri AwsLexIntent#uri}.
	// Experimental.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

