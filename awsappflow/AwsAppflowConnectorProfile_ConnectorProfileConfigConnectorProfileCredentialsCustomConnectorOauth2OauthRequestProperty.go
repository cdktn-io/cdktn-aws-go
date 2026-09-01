package awsappflow


// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequestProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#auth_code AwsAppflowConnectorProfile#auth_code}.
	// Experimental.
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#redirect_uri AwsAppflowConnectorProfile#redirect_uri}.
	// Experimental.
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

