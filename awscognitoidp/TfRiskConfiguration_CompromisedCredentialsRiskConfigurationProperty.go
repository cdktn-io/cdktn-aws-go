package awscognitoidp


// Experimental.
type TfRiskConfiguration_CompromisedCredentialsRiskConfigurationProperty struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#actions TfRiskConfiguration#actions}
	// Experimental.
	Actions *TfRiskConfiguration_CompromisedCredentialsRiskConfigurationActionsProperty `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#event_filter TfRiskConfiguration#event_filter}.
	// Experimental.
	EventFilter *[]*string `field:"optional" json:"eventFilter" yaml:"eventFilter"`
}

