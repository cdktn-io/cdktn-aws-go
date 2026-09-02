package awslexmodelbuilding


// Experimental.
type TfBot_IntentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#intent_name TfBot#intent_name}.
	// Experimental.
	IntentName *string `field:"required" json:"intentName" yaml:"intentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot#intent_version TfBot#intent_version}.
	// Experimental.
	IntentVersion *string `field:"required" json:"intentVersion" yaml:"intentVersion"`
}

