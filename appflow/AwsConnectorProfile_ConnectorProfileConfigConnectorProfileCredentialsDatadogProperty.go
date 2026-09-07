package appflow


// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#api_key AwsConnectorProfile#api_key}.
	// Experimental.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#application_key AwsConnectorProfile#application_key}.
	// Experimental.
	ApplicationKey *string `field:"required" json:"applicationKey" yaml:"applicationKey"`
}

