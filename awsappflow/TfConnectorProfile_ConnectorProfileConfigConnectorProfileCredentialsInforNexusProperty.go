package awsappflow


// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#access_key_id TfConnectorProfile#access_key_id}.
	// Experimental.
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#datakey TfConnectorProfile#datakey}.
	// Experimental.
	Datakey *string `field:"required" json:"datakey" yaml:"datakey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#secret_access_key TfConnectorProfile#secret_access_key}.
	// Experimental.
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#user_id TfConnectorProfile#user_id}.
	// Experimental.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

