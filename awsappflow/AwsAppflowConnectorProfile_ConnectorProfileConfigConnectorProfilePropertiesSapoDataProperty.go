package awsappflow


// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesSapoDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#application_host_url AwsAppflowConnectorProfile#application_host_url}.
	// Experimental.
	ApplicationHostUrl *string `field:"required" json:"applicationHostUrl" yaml:"applicationHostUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#application_service_path AwsAppflowConnectorProfile#application_service_path}.
	// Experimental.
	ApplicationServicePath *string `field:"required" json:"applicationServicePath" yaml:"applicationServicePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#client_number AwsAppflowConnectorProfile#client_number}.
	// Experimental.
	ClientNumber *string `field:"required" json:"clientNumber" yaml:"clientNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#port_number AwsAppflowConnectorProfile#port_number}.
	// Experimental.
	PortNumber *float64 `field:"required" json:"portNumber" yaml:"portNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#logon_language AwsAppflowConnectorProfile#logon_language}.
	// Experimental.
	LogonLanguage *string `field:"optional" json:"logonLanguage" yaml:"logonLanguage"`
	// oauth_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth_properties AwsAppflowConnectorProfile#oauth_properties}
	// Experimental.
	OauthProperties *AwsAppflowConnectorProfile_OauthPropertiesProperty `field:"optional" json:"oauthProperties" yaml:"oauthProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#private_link_service_name AwsAppflowConnectorProfile#private_link_service_name}.
	// Experimental.
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
}

