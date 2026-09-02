package awsappflow


// Experimental.
type TfConnectorProfile_Oauth2PropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth2_grant_type TfConnectorProfile#oauth2_grant_type}.
	// Experimental.
	Oauth2GrantType *string `field:"required" json:"oauth2GrantType" yaml:"oauth2GrantType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#token_url TfConnectorProfile#token_url}.
	// Experimental.
	TokenUrl *string `field:"required" json:"tokenUrl" yaml:"tokenUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#token_url_custom_properties TfConnectorProfile#token_url_custom_properties}.
	// Experimental.
	TokenUrlCustomProperties *map[string]*string `field:"optional" json:"tokenUrlCustomProperties" yaml:"tokenUrlCustomProperties"`
}

