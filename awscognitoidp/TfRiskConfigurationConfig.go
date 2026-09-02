package awscognitoidp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRiskConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#user_pool_id TfRiskConfiguration#user_pool_id}.
	// Experimental.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// account_takeover_risk_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#account_takeover_risk_configuration TfRiskConfiguration#account_takeover_risk_configuration}
	// Experimental.
	AccountTakeoverRiskConfiguration *TfRiskConfiguration_AccountTakeoverRiskConfigurationProperty `field:"optional" json:"accountTakeoverRiskConfiguration" yaml:"accountTakeoverRiskConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#client_id TfRiskConfiguration#client_id}.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// compromised_credentials_risk_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#compromised_credentials_risk_configuration TfRiskConfiguration#compromised_credentials_risk_configuration}
	// Experimental.
	CompromisedCredentialsRiskConfiguration *TfRiskConfiguration_CompromisedCredentialsRiskConfigurationProperty `field:"optional" json:"compromisedCredentialsRiskConfiguration" yaml:"compromisedCredentialsRiskConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#id TfRiskConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#region TfRiskConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// risk_exception_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_risk_configuration#risk_exception_configuration TfRiskConfiguration#risk_exception_configuration}
	// Experimental.
	RiskExceptionConfiguration *TfRiskConfiguration_RiskExceptionConfigurationProperty `field:"optional" json:"riskExceptionConfiguration" yaml:"riskExceptionConfiguration"`
}

