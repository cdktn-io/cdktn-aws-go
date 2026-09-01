package awsappflow


// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty struct {
	// basic_auth_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#basic_auth_credentials AwsAppflowConnectorProfile#basic_auth_credentials}
	// Experimental.
	BasicAuthCredentials *AwsAppflowConnectorProfile_BasicAuthCredentialsProperty `field:"optional" json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// oauth_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth_credentials AwsAppflowConnectorProfile#oauth_credentials}
	// Experimental.
	OauthCredentials *AwsAppflowConnectorProfile_OauthCredentialsProperty `field:"optional" json:"oauthCredentials" yaml:"oauthCredentials"`
}

