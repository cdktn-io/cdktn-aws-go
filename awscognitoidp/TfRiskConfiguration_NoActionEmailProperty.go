package awscognitoidp


// Experimental.
type TfRiskConfiguration_NoActionEmailProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#html_body TfRiskConfiguration#html_body}.
	// Experimental.
	HtmlBody *string `field:"required" json:"htmlBody" yaml:"htmlBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#subject TfRiskConfiguration#subject}.
	// Experimental.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#text_body TfRiskConfiguration#text_body}.
	// Experimental.
	TextBody *string `field:"required" json:"textBody" yaml:"textBody"`
}

