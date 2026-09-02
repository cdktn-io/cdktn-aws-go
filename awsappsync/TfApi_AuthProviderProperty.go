package awsappsync


// Experimental.
type TfApi_AuthProviderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#auth_type TfApi#auth_type}.
	// Experimental.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// cognito_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#cognito_config TfApi#cognito_config}
	// Experimental.
	CognitoConfig interface{} `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#lambda_authorizer_config TfApi#lambda_authorizer_config}
	// Experimental.
	LambdaAuthorizerConfig interface{} `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#openid_connect_config TfApi#openid_connect_config}
	// Experimental.
	OpenidConnectConfig interface{} `field:"optional" json:"openidConnectConfig" yaml:"openidConnectConfig"`
}

