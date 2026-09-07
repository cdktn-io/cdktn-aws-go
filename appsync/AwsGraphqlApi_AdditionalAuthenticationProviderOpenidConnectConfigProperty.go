package appsync


// Experimental.
type AwsGraphqlApi_AdditionalAuthenticationProviderOpenidConnectConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#issuer AwsGraphqlApi#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#auth_ttl AwsGraphqlApi#auth_ttl}.
	// Experimental.
	AuthTtl *float64 `field:"optional" json:"authTtl" yaml:"authTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#client_id AwsGraphqlApi#client_id}.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#iat_ttl AwsGraphqlApi#iat_ttl}.
	// Experimental.
	IatTtl *float64 `field:"optional" json:"iatTtl" yaml:"iatTtl"`
}

