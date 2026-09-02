package awsappflow


// Experimental.
type TfConnectorProfile_Oauth2Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#access_token TfConnectorProfile#access_token}.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#client_id TfConnectorProfile#client_id}.
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#client_secret TfConnectorProfile#client_secret}.
	// Experimental.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// oauth_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth_request TfConnectorProfile#oauth_request}
	// Experimental.
	OauthRequest *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestProperty `field:"optional" json:"oauthRequest" yaml:"oauthRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#refresh_token TfConnectorProfile#refresh_token}.
	// Experimental.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

