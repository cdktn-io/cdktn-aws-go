package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#application_host_url TfConnectorProfile#application_host_url}.
	// Experimental.
	ApplicationHostUrl *string `field:"required" json:"applicationHostUrl" yaml:"applicationHostUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#application_service_path TfConnectorProfile#application_service_path}.
	// Experimental.
	ApplicationServicePath *string `field:"required" json:"applicationServicePath" yaml:"applicationServicePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#client_number TfConnectorProfile#client_number}.
	// Experimental.
	ClientNumber *string `field:"required" json:"clientNumber" yaml:"clientNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#port_number TfConnectorProfile#port_number}.
	// Experimental.
	PortNumber *float64 `field:"required" json:"portNumber" yaml:"portNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#logon_language TfConnectorProfile#logon_language}.
	// Experimental.
	LogonLanguage *string `field:"optional" json:"logonLanguage" yaml:"logonLanguage"`
	// oauth_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth_properties TfConnectorProfile#oauth_properties}
	// Experimental.
	OauthProperties *TfConnectorProfile_OauthPropertiesProperty `field:"optional" json:"oauthProperties" yaml:"oauthProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#private_link_service_name TfConnectorProfile#private_link_service_name}.
	// Experimental.
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
}

