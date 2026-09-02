package awscognitoidp


// Experimental.
type TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty struct {
	// high_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#high_action TfRiskConfiguration#high_action}
	// Experimental.
	HighAction *TfRiskConfiguration_HighActionProperty `field:"optional" json:"highAction" yaml:"highAction"`
	// low_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#low_action TfRiskConfiguration#low_action}
	// Experimental.
	LowAction *TfRiskConfiguration_LowActionProperty `field:"optional" json:"lowAction" yaml:"lowAction"`
	// medium_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#medium_action TfRiskConfiguration#medium_action}
	// Experimental.
	MediumAction *TfRiskConfiguration_MediumActionProperty `field:"optional" json:"mediumAction" yaml:"mediumAction"`
}

