package awscognitoidp


// Experimental.
type AwsCognitoRiskConfiguration_MediumActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#event_action AwsCognitoRiskConfiguration#event_action}.
	// Experimental.
	EventAction *string `field:"required" json:"eventAction" yaml:"eventAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#notify AwsCognitoRiskConfiguration#notify}.
	// Experimental.
	Notify interface{} `field:"required" json:"notify" yaml:"notify"`
}

