package appsync


// Experimental.
type AwsGraphqlApi_AdditionalAuthenticationProviderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#authentication_type AwsGraphqlApi#authentication_type}.
	// Experimental.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#lambda_authorizer_config AwsGraphqlApi#lambda_authorizer_config}
	// Experimental.
	LambdaAuthorizerConfig *AwsGraphqlApi_AdditionalAuthenticationProviderLambdaAuthorizerConfigProperty `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#openid_connect_config AwsGraphqlApi#openid_connect_config}
	// Experimental.
	OpenidConnectConfig *AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty `field:"optional" json:"openidConnectConfig" yaml:"openidConnectConfig"`
	// user_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#user_pool_config AwsGraphqlApi#user_pool_config}
	// Experimental.
	UserPoolConfig *AwsGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
}

