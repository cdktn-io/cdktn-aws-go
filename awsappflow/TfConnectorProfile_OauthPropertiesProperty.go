package awsappflow


// Experimental.
type TfConnectorProfile_OauthPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#auth_code_url TfConnectorProfile#auth_code_url}.
	// Experimental.
	AuthCodeUrl *string `field:"required" json:"authCodeUrl" yaml:"authCodeUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth_scopes TfConnectorProfile#oauth_scopes}.
	// Experimental.
	OauthScopes *[]*string `field:"required" json:"oauthScopes" yaml:"oauthScopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#token_url TfConnectorProfile#token_url}.
	// Experimental.
	TokenUrl *string `field:"required" json:"tokenUrl" yaml:"tokenUrl"`
}

