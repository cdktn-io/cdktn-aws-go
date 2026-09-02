package awsappsync


// Experimental.
type TfGraphqlApi_AdditionalAuthenticationProviderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#authentication_type TfGraphqlApi#authentication_type}.
	// Experimental.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#lambda_authorizer_config TfGraphqlApi#lambda_authorizer_config}
	// Experimental.
	LambdaAuthorizerConfig *TfGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#openid_connect_config TfGraphqlApi#openid_connect_config}
	// Experimental.
	OpenidConnectConfig *TfGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty `field:"optional" json:"openidConnectConfig" yaml:"openidConnectConfig"`
	// user_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#user_pool_config TfGraphqlApi#user_pool_config}
	// Experimental.
	UserPoolConfig *TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
}

