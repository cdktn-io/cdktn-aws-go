package appflow


// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty struct {
	// oauth2_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth2_properties AwsConnectorProfile#oauth2_properties}
	// Experimental.
	Oauth2Properties *AwsConnectorProfile_Oauth2PropertiesProperty `field:"optional" json:"oauth2Properties" yaml:"oauth2Properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#profile_properties AwsConnectorProfile#profile_properties}.
	// Experimental.
	ProfileProperties *map[string]*string `field:"optional" json:"profileProperties" yaml:"profileProperties"`
}

