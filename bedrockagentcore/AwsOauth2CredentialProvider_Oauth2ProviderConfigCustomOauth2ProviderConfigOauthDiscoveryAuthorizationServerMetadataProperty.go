package bedrockagentcore


// Experimental.
type AwsOauth2CredentialProvider_Oauth2ProviderConfigCustomOauth2ProviderConfigOauthDiscoveryAuthorizationServerMetadataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#authorization_endpoint AwsOauth2CredentialProvider#authorization_endpoint}.
	// Experimental.
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#issuer AwsOauth2CredentialProvider#issuer}.
	// Experimental.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#token_endpoint AwsOauth2CredentialProvider#token_endpoint}.
	// Experimental.
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#response_types AwsOauth2CredentialProvider#response_types}.
	// Experimental.
	ResponseTypes *[]*string `field:"optional" json:"responseTypes" yaml:"responseTypes"`
}

