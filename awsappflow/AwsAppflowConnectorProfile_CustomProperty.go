package awsappflow


// Experimental.
type AwsAppflowConnectorProfile_CustomProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#custom_authentication_type AwsAppflowConnectorProfile#custom_authentication_type}.
	// Experimental.
	CustomAuthenticationType *string `field:"required" json:"customAuthenticationType" yaml:"customAuthenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#credentials_map AwsAppflowConnectorProfile#credentials_map}.
	// Experimental.
	CredentialsMap *map[string]*string `field:"optional" json:"credentialsMap" yaml:"credentialsMap"`
}

