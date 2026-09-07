package cognitoidp


// Experimental.
type AwsRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty struct {
	// high_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#high_action AwsRiskConfiguration#high_action}
	// Experimental.
	HighAction *AwsRiskConfiguration_HighActionProperty `field:"optional" json:"highAction" yaml:"highAction"`
	// low_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#low_action AwsRiskConfiguration#low_action}
	// Experimental.
	LowAction *AwsRiskConfiguration_LowActionProperty `field:"optional" json:"lowAction" yaml:"lowAction"`
	// medium_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#medium_action AwsRiskConfiguration#medium_action}
	// Experimental.
	MediumAction *AwsRiskConfiguration_MediumActionProperty `field:"optional" json:"mediumAction" yaml:"mediumAction"`
}

