package awsappsync


// Experimental.
type TfGraphqlApi_AdditionalAuthenticationProviderUserPoolConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#user_pool_id TfGraphqlApi#user_pool_id}.
	// Experimental.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#app_id_client_regex TfGraphqlApi#app_id_client_regex}.
	// Experimental.
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#aws_region TfGraphqlApi#aws_region}.
	// Experimental.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
}

