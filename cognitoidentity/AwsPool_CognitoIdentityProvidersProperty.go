package cognitoidentity


// Experimental.
type AwsPool_CognitoIdentityProvidersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool#client_id AwsPool#client_id}.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool#provider_name AwsPool#provider_name}.
	// Experimental.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_identity_pool#server_side_token_check AwsPool#server_side_token_check}.
	// Experimental.
	ServerSideTokenCheck interface{} `field:"optional" json:"serverSideTokenCheck" yaml:"serverSideTokenCheck"`
}

