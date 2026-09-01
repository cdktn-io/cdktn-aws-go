package awsappsync


// Experimental.
type AwsAppsyncGraphqlApi_OpenidConnectConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#issuer AwsAppsyncGraphqlApi#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#auth_ttl AwsAppsyncGraphqlApi#auth_ttl}.
	// Experimental.
	AuthTtl *float64 `field:"optional" json:"authTtl" yaml:"authTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#client_id AwsAppsyncGraphqlApi#client_id}.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#iat_ttl AwsAppsyncGraphqlApi#iat_ttl}.
	// Experimental.
	IatTtl *float64 `field:"optional" json:"iatTtl" yaml:"iatTtl"`
}

