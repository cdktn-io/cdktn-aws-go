package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#client_id TfConnectorProfile#client_id}.
	// Experimental.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#client_secret TfConnectorProfile#client_secret}.
	// Experimental.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#access_token TfConnectorProfile#access_token}.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// oauth_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth_request TfConnectorProfile#oauth_request}
	// Experimental.
	OauthRequest *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackOauthRequestProperty `field:"optional" json:"oauthRequest" yaml:"oauthRequest"`
}

