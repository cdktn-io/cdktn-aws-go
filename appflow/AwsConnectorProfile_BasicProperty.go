package appflow


// Experimental.
type AwsConnectorProfile_BasicProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#password AwsConnectorProfile#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#username AwsConnectorProfile#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

