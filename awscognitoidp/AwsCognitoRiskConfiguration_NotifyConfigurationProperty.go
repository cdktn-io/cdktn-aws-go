package awscognitoidp


// Experimental.
type AwsCognitoRiskConfiguration_NotifyConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#source_arn AwsCognitoRiskConfiguration#source_arn}.
	// Experimental.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// block_email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#block_email AwsCognitoRiskConfiguration#block_email}
	// Experimental.
	BlockEmail *AwsCognitoRiskConfiguration_BlockEmailProperty `field:"optional" json:"blockEmail" yaml:"blockEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#from AwsCognitoRiskConfiguration#from}.
	// Experimental.
	From *string `field:"optional" json:"from" yaml:"from"`
	// mfa_email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#mfa_email AwsCognitoRiskConfiguration#mfa_email}
	// Experimental.
	MfaEmail *AwsCognitoRiskConfiguration_MfaEmailProperty `field:"optional" json:"mfaEmail" yaml:"mfaEmail"`
	// no_action_email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#no_action_email AwsCognitoRiskConfiguration#no_action_email}
	// Experimental.
	NoActionEmail *AwsCognitoRiskConfiguration_NoActionEmailProperty `field:"optional" json:"noActionEmail" yaml:"noActionEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#reply_to AwsCognitoRiskConfiguration#reply_to}.
	// Experimental.
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
}

