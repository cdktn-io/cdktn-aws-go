package awscognitoidp


// Experimental.
type TfRiskConfiguration_HighActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#event_action TfRiskConfiguration#event_action}.
	// Experimental.
	EventAction *string `field:"required" json:"eventAction" yaml:"eventAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#notify TfRiskConfiguration#notify}.
	// Experimental.
	Notify interface{} `field:"required" json:"notify" yaml:"notify"`
}

