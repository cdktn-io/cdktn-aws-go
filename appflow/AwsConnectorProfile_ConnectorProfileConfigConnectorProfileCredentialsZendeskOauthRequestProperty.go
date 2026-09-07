package appflow


// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskOauthRequestProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#auth_code AwsConnectorProfile#auth_code}.
	// Experimental.
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#redirect_uri AwsConnectorProfile#redirect_uri}.
	// Experimental.
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

