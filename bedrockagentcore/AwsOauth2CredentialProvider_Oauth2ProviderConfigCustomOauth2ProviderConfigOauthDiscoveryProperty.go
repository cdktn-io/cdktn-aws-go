package bedrockagentcore


// Experimental.
type AwsOauth2CredentialProvider_Oauth2ProviderConfigCustomOauth2ProviderConfigOauthDiscoveryProperty struct {
	// authorization_server_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#authorization_server_metadata AwsOauth2CredentialProvider#authorization_server_metadata}
	// Experimental.
	AuthorizationServerMetadata interface{} `field:"optional" json:"authorizationServerMetadata" yaml:"authorizationServerMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#discovery_url AwsOauth2CredentialProvider#discovery_url}.
	// Experimental.
	DiscoveryUrl *string `field:"optional" json:"discoveryUrl" yaml:"discoveryUrl"`
}

