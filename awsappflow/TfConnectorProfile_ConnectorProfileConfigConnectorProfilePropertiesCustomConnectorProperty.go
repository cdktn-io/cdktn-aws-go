package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfilePropertiesCustomConnectorProperty struct {
	// oauth2_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth2_properties TfConnectorProfile#oauth2_properties}
	// Experimental.
	Oauth2Properties *TfConnectorProfile_Oauth2PropertiesProperty `field:"optional" json:"oauth2Properties" yaml:"oauth2Properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#profile_properties TfConnectorProfile#profile_properties}.
	// Experimental.
	ProfileProperties *map[string]*string `field:"optional" json:"profileProperties" yaml:"profileProperties"`
}

