package appflow


// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#authentication_type AwsConnectorProfile#authentication_type}.
	// Experimental.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#api_key AwsConnectorProfile#api_key}
	// Experimental.
	ApiKey *AwsConnectorProfile_ApiKeyProperty `field:"optional" json:"apiKey" yaml:"apiKey"`
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#basic AwsConnectorProfile#basic}
	// Experimental.
	Basic *AwsConnectorProfile_BasicProperty `field:"optional" json:"basic" yaml:"basic"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#custom AwsConnectorProfile#custom}
	// Experimental.
	Custom *AwsConnectorProfile_CustomProperty `field:"optional" json:"custom" yaml:"custom"`
	// oauth2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth2 AwsConnectorProfile#oauth2}
	// Experimental.
	Oauth2 *AwsConnectorProfile_Oauth2Property `field:"optional" json:"oauth2" yaml:"oauth2"`
}

