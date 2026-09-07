package cognitoidp


// Experimental.
type AwsRiskConfiguration_NotifyConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#source_arn AwsRiskConfiguration#source_arn}.
	// Experimental.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// block_email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#block_email AwsRiskConfiguration#block_email}
	// Experimental.
	BlockEmail *AwsRiskConfiguration_BlockEmailProperty `field:"optional" json:"blockEmail" yaml:"blockEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#from AwsRiskConfiguration#from}.
	// Experimental.
	From *string `field:"optional" json:"from" yaml:"from"`
	// mfa_email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#mfa_email AwsRiskConfiguration#mfa_email}
	// Experimental.
	MfaEmail *AwsRiskConfiguration_MfaEmailProperty `field:"optional" json:"mfaEmail" yaml:"mfaEmail"`
	// no_action_email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#no_action_email AwsRiskConfiguration#no_action_email}
	// Experimental.
	NoActionEmail *AwsRiskConfiguration_NoActionEmailProperty `field:"optional" json:"noActionEmail" yaml:"noActionEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#reply_to AwsRiskConfiguration#reply_to}.
	// Experimental.
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
}

