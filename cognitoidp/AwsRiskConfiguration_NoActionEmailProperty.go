package cognitoidp


// Experimental.
type AwsRiskConfiguration_NoActionEmailProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#html_body AwsRiskConfiguration#html_body}.
	// Experimental.
	HtmlBody *string `field:"required" json:"htmlBody" yaml:"htmlBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#subject AwsRiskConfiguration#subject}.
	// Experimental.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#text_body AwsRiskConfiguration#text_body}.
	// Experimental.
	TextBody *string `field:"required" json:"textBody" yaml:"textBody"`
}

