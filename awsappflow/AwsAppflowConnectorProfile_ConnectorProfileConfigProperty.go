package awsappflow


// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileConfigProperty struct {
	// connector_profile_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#connector_profile_credentials AwsAppflowConnectorProfile#connector_profile_credentials}
	// Experimental.
	ConnectorProfileCredentials *AwsAppflowConnectorProfile_ConnectorProfileCredentialsProperty `field:"required" json:"connectorProfileCredentials" yaml:"connectorProfileCredentials"`
	// connector_profile_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#connector_profile_properties AwsAppflowConnectorProfile#connector_profile_properties}
	// Experimental.
	ConnectorProfileProperties *AwsAppflowConnectorProfile_ConnectorProfilePropertiesProperty `field:"required" json:"connectorProfileProperties" yaml:"connectorProfileProperties"`
}

