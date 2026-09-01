package awsappflow


// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#authentication_type AwsAppflowConnectorProfile#authentication_type}.
	// Experimental.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#api_key AwsAppflowConnectorProfile#api_key}
	// Experimental.
	ApiKey *AwsAppflowConnectorProfile_ApiKeyProperty `field:"optional" json:"apiKey" yaml:"apiKey"`
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#basic AwsAppflowConnectorProfile#basic}
	// Experimental.
	Basic *AwsAppflowConnectorProfile_BasicProperty `field:"optional" json:"basic" yaml:"basic"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#custom AwsAppflowConnectorProfile#custom}
	// Experimental.
	Custom *AwsAppflowConnectorProfile_CustomProperty `field:"optional" json:"custom" yaml:"custom"`
	// oauth2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#oauth2 AwsAppflowConnectorProfile#oauth2}
	// Experimental.
	Oauth2 *AwsAppflowConnectorProfile_Oauth2Property `field:"optional" json:"oauth2" yaml:"oauth2"`
}

