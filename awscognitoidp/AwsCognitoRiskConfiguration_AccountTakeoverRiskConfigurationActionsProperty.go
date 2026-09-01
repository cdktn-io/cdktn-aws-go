package awscognitoidp


// Experimental.
type AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty struct {
	// high_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#high_action AwsCognitoRiskConfiguration#high_action}
	// Experimental.
	HighAction *AwsCognitoRiskConfiguration_HighActionProperty `field:"optional" json:"highAction" yaml:"highAction"`
	// low_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#low_action AwsCognitoRiskConfiguration#low_action}
	// Experimental.
	LowAction *AwsCognitoRiskConfiguration_LowActionProperty `field:"optional" json:"lowAction" yaml:"lowAction"`
	// medium_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#medium_action AwsCognitoRiskConfiguration#medium_action}
	// Experimental.
	MediumAction *AwsCognitoRiskConfiguration_MediumActionProperty `field:"optional" json:"mediumAction" yaml:"mediumAction"`
}

