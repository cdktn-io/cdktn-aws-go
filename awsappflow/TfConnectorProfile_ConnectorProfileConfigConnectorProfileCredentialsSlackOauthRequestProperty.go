package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackOauthRequestProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#auth_code TfConnectorProfile#auth_code}.
	// Experimental.
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#redirect_uri TfConnectorProfile#redirect_uri}.
	// Experimental.
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

