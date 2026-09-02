package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#access_token TfConnectorProfile#access_token}.
	// Experimental.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#client_credentials_arn TfConnectorProfile#client_credentials_arn}.
	// Experimental.
	ClientCredentialsArn *string `field:"optional" json:"clientCredentialsArn" yaml:"clientCredentialsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#jwt_token TfConnectorProfile#jwt_token}.
	// Experimental.
	JwtToken *string `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth2_grant_type TfConnectorProfile#oauth2_grant_type}.
	// Experimental.
	Oauth2GrantType *string `field:"optional" json:"oauth2GrantType" yaml:"oauth2GrantType"`
	// oauth_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth_request TfConnectorProfile#oauth_request}
	// Experimental.
	OauthRequest *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequestProperty `field:"optional" json:"oauthRequest" yaml:"oauthRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#refresh_token TfConnectorProfile#refresh_token}.
	// Experimental.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

