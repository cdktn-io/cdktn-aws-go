package awscognitoidp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserPoolClientConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#name TfUserPoolClient#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#user_pool_id TfUserPoolClient#user_pool_id}.
	// Experimental.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#access_token_validity TfUserPoolClient#access_token_validity}.
	// Experimental.
	AccessTokenValidity *float64 `field:"optional" json:"accessTokenValidity" yaml:"accessTokenValidity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#allowed_oauth_flows TfUserPoolClient#allowed_oauth_flows}.
	// Experimental.
	AllowedOauthFlows *[]*string `field:"optional" json:"allowedOauthFlows" yaml:"allowedOauthFlows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#allowed_oauth_flows_user_pool_client TfUserPoolClient#allowed_oauth_flows_user_pool_client}.
	// Experimental.
	AllowedOauthFlowsUserPoolClient interface{} `field:"optional" json:"allowedOauthFlowsUserPoolClient" yaml:"allowedOauthFlowsUserPoolClient"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#allowed_oauth_scopes TfUserPoolClient#allowed_oauth_scopes}.
	// Experimental.
	AllowedOauthScopes *[]*string `field:"optional" json:"allowedOauthScopes" yaml:"allowedOauthScopes"`
	// analytics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#analytics_configuration TfUserPoolClient#analytics_configuration}
	// Experimental.
	AnalyticsConfiguration interface{} `field:"optional" json:"analyticsConfiguration" yaml:"analyticsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#auth_session_validity TfUserPoolClient#auth_session_validity}.
	// Experimental.
	AuthSessionValidity *float64 `field:"optional" json:"authSessionValidity" yaml:"authSessionValidity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#callback_urls TfUserPoolClient#callback_urls}.
	// Experimental.
	CallbackUrls *[]*string `field:"optional" json:"callbackUrls" yaml:"callbackUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#default_redirect_uri TfUserPoolClient#default_redirect_uri}.
	// Experimental.
	DefaultRedirectUri *string `field:"optional" json:"defaultRedirectUri" yaml:"defaultRedirectUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#enable_propagate_additional_user_context_data TfUserPoolClient#enable_propagate_additional_user_context_data}.
	// Experimental.
	EnablePropagateAdditionalUserContextData interface{} `field:"optional" json:"enablePropagateAdditionalUserContextData" yaml:"enablePropagateAdditionalUserContextData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#enable_token_revocation TfUserPoolClient#enable_token_revocation}.
	// Experimental.
	EnableTokenRevocation interface{} `field:"optional" json:"enableTokenRevocation" yaml:"enableTokenRevocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#explicit_auth_flows TfUserPoolClient#explicit_auth_flows}.
	// Experimental.
	ExplicitAuthFlows *[]*string `field:"optional" json:"explicitAuthFlows" yaml:"explicitAuthFlows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#generate_secret TfUserPoolClient#generate_secret}.
	// Experimental.
	GenerateSecret interface{} `field:"optional" json:"generateSecret" yaml:"generateSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#id_token_validity TfUserPoolClient#id_token_validity}.
	// Experimental.
	IdTokenValidity *float64 `field:"optional" json:"idTokenValidity" yaml:"idTokenValidity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#logout_urls TfUserPoolClient#logout_urls}.
	// Experimental.
	LogoutUrls *[]*string `field:"optional" json:"logoutUrls" yaml:"logoutUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#prevent_user_existence_errors TfUserPoolClient#prevent_user_existence_errors}.
	// Experimental.
	PreventUserExistenceErrors *string `field:"optional" json:"preventUserExistenceErrors" yaml:"preventUserExistenceErrors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#read_attributes TfUserPoolClient#read_attributes}.
	// Experimental.
	ReadAttributes *[]*string `field:"optional" json:"readAttributes" yaml:"readAttributes"`
	// refresh_token_rotation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#refresh_token_rotation TfUserPoolClient#refresh_token_rotation}
	// Experimental.
	RefreshTokenRotation interface{} `field:"optional" json:"refreshTokenRotation" yaml:"refreshTokenRotation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#refresh_token_validity TfUserPoolClient#refresh_token_validity}.
	// Experimental.
	RefreshTokenValidity *float64 `field:"optional" json:"refreshTokenValidity" yaml:"refreshTokenValidity"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#region TfUserPoolClient#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#supported_identity_providers TfUserPoolClient#supported_identity_providers}.
	// Experimental.
	SupportedIdentityProviders *[]*string `field:"optional" json:"supportedIdentityProviders" yaml:"supportedIdentityProviders"`
	// token_validity_units block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#token_validity_units TfUserPoolClient#token_validity_units}
	// Experimental.
	TokenValidityUnits interface{} `field:"optional" json:"tokenValidityUnits" yaml:"tokenValidityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool_client#write_attributes TfUserPoolClient#write_attributes}.
	// Experimental.
	WriteAttributes *[]*string `field:"optional" json:"writeAttributes" yaml:"writeAttributes"`
}

