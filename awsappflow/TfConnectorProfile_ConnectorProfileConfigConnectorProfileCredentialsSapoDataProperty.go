package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty struct {
	// basic_auth_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#basic_auth_credentials TfConnectorProfile#basic_auth_credentials}
	// Experimental.
	BasicAuthCredentials *TfConnectorProfile_BasicAuthCredentialsProperty `field:"optional" json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// oauth_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth_credentials TfConnectorProfile#oauth_credentials}
	// Experimental.
	OauthCredentials *TfConnectorProfile_OauthCredentialsProperty `field:"optional" json:"oauthCredentials" yaml:"oauthCredentials"`
}

