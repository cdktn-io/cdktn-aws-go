package appflow


// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigProperty struct {
	// connector_profile_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#connector_profile_credentials AwsConnectorProfile#connector_profile_credentials}
	// Experimental.
	ConnectorProfileCredentials *AwsConnectorProfile_ConnectorProfileCredentialsProperty `field:"required" json:"connectorProfileCredentials" yaml:"connectorProfileCredentials"`
	// connector_profile_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#connector_profile_properties AwsConnectorProfile#connector_profile_properties}
	// Experimental.
	ConnectorProfileProperties *AwsConnectorProfile_ConnectorProfilePropertiesProperty `field:"required" json:"connectorProfileProperties" yaml:"connectorProfileProperties"`
}

