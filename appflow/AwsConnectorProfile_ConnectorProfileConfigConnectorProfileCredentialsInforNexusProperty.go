package appflow


// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#access_key_id AwsConnectorProfile#access_key_id}.
	// Experimental.
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#datakey AwsConnectorProfile#datakey}.
	// Experimental.
	Datakey *string `field:"required" json:"datakey" yaml:"datakey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#secret_access_key AwsConnectorProfile#secret_access_key}.
	// Experimental.
	SecretAccessKey *string `field:"required" json:"secretAccessKey" yaml:"secretAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_connector_profile#user_id AwsConnectorProfile#user_id}.
	// Experimental.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

