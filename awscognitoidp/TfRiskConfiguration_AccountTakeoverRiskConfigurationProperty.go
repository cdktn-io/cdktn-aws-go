package awscognitoidp


// Experimental.
type TfRiskConfiguration_AccountTakeoverRiskConfigurationProperty struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#actions TfRiskConfiguration#actions}
	// Experimental.
	Actions *TfRiskConfiguration_AccountTakeoverRiskConfigurationActionsProperty `field:"required" json:"actions" yaml:"actions"`
	// notify_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#notify_configuration TfRiskConfiguration#notify_configuration}
	// Experimental.
	NotifyConfiguration *TfRiskConfiguration_NotifyConfigurationProperty `field:"optional" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

