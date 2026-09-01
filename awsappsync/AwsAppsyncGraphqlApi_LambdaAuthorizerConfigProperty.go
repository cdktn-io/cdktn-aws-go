package awsappsync


// Experimental.
type AwsAppsyncGraphqlApi_LambdaAuthorizerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#authorizer_uri AwsAppsyncGraphqlApi#authorizer_uri}.
	// Experimental.
	AuthorizerUri *string `field:"required" json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#authorizer_result_ttl_in_seconds AwsAppsyncGraphqlApi#authorizer_result_ttl_in_seconds}.
	// Experimental.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#identity_validation_expression AwsAppsyncGraphqlApi#identity_validation_expression}.
	// Experimental.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
}

