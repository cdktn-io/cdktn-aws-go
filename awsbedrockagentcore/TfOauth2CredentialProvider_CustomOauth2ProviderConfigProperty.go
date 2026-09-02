package awsbedrockagentcore


// Experimental.
type TfOauth2CredentialProvider_CustomOauth2ProviderConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#client_credentials_wo_version TfOauth2CredentialProvider#client_credentials_wo_version}.
	// Experimental.
	ClientCredentialsWoVersion *float64 `field:"optional" json:"clientCredentialsWoVersion" yaml:"clientCredentialsWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#client_id TfOauth2CredentialProvider#client_id}.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#client_id_wo TfOauth2CredentialProvider#client_id_wo}.
	// Experimental.
	ClientIdWo *string `field:"optional" json:"clientIdWo" yaml:"clientIdWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#client_secret TfOauth2CredentialProvider#client_secret}.
	// Experimental.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#client_secret_wo TfOauth2CredentialProvider#client_secret_wo}.
	// Experimental.
	ClientSecretWo *string `field:"optional" json:"clientSecretWo" yaml:"clientSecretWo"`
	// oauth_discovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_oauth2_credential_provider#oauth_discovery TfOauth2CredentialProvider#oauth_discovery}
	// Experimental.
	OauthDiscovery interface{} `field:"optional" json:"oauthDiscovery" yaml:"oauthDiscovery"`
}

