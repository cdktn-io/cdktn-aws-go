package cognitoidp


// Experimental.
type AwsRiskConfiguration_AccountTakeoverRiskConfigurationProperty struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#actions AwsRiskConfiguration#actions}
	// Experimental.
	Actions *AwsRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty `field:"required" json:"actions" yaml:"actions"`
	// notify_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#notify_configuration AwsRiskConfiguration#notify_configuration}
	// Experimental.
	NotifyConfiguration *AwsRiskConfiguration_NotifyConfigurationProperty `field:"optional" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

