package appsync


// Experimental.
type AwsApi_LambdaAuthorizerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#authorizer_uri AwsApi#authorizer_uri}.
	// Experimental.
	AuthorizerUri *string `field:"required" json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#authorizer_result_ttl_in_seconds AwsApi#authorizer_result_ttl_in_seconds}.
	// Experimental.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#identity_validation_expression AwsApi#identity_validation_expression}.
	// Experimental.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
}

