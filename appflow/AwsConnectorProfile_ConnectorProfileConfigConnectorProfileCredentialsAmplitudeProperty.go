package appflow


// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#api_key AwsConnectorProfile#api_key}.
	// Experimental.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#secret_key AwsConnectorProfile#secret_key}.
	// Experimental.
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
}

