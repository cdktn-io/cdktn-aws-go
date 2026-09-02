package awsappflow


// Experimental.
type TfConnectorProfile_ApiKeyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#api_key TfConnectorProfile#api_key}.
	// Experimental.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#api_secret_key TfConnectorProfile#api_secret_key}.
	// Experimental.
	ApiSecretKey *string `field:"optional" json:"apiSecretKey" yaml:"apiSecretKey"`
}

