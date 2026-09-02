package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#password TfConnectorProfile#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#username TfConnectorProfile#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

