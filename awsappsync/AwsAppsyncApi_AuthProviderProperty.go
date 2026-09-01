package awsappsync


// Experimental.
type AwsAppsyncApi_AuthProviderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#auth_type AwsAppsyncApi#auth_type}.
	// Experimental.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// cognito_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#cognito_config AwsAppsyncApi#cognito_config}
	// Experimental.
	CognitoConfig interface{} `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#lambda_authorizer_config AwsAppsyncApi#lambda_authorizer_config}
	// Experimental.
	LambdaAuthorizerConfig interface{} `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#openid_connect_config AwsAppsyncApi#openid_connect_config}
	// Experimental.
	OpenidConnectConfig interface{} `field:"optional" json:"openidConnectConfig" yaml:"openidConnectConfig"`
}

