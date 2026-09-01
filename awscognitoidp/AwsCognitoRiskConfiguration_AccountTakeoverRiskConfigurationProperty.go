package awscognitoidp


// Experimental.
type AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationProperty struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#actions AwsCognitoRiskConfiguration#actions}
	// Experimental.
	Actions *AwsCognitoRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty `field:"required" json:"actions" yaml:"actions"`
	// notify_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#notify_configuration AwsCognitoRiskConfiguration#notify_configuration}
	// Experimental.
	NotifyConfiguration *AwsCognitoRiskConfiguration_NotifyConfigurationProperty `field:"optional" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

