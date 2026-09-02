package awsbedrockagents


// Experimental.
type TfDataSource_DataSourceConfigurationConfluenceConfigurationSourceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#auth_type TfDataSource#auth_type}.
	// Experimental.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#credentials_secret_arn TfDataSource#credentials_secret_arn}.
	// Experimental.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#host_type TfDataSource#host_type}.
	// Experimental.
	HostType *string `field:"required" json:"hostType" yaml:"hostType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#host_url TfDataSource#host_url}.
	// Experimental.
	HostUrl *string `field:"required" json:"hostUrl" yaml:"hostUrl"`
}

