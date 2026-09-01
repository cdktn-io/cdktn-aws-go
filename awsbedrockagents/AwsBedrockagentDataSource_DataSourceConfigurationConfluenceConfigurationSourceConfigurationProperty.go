package awsbedrockagents


// Experimental.
type AwsBedrockagentDataSource_DataSourceConfigurationConfluenceConfigurationSourceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#auth_type AwsBedrockagentDataSource#auth_type}.
	// Experimental.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#credentials_secret_arn AwsBedrockagentDataSource#credentials_secret_arn}.
	// Experimental.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#host_type AwsBedrockagentDataSource#host_type}.
	// Experimental.
	HostType *string `field:"required" json:"hostType" yaml:"hostType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#host_url AwsBedrockagentDataSource#host_url}.
	// Experimental.
	HostUrl *string `field:"required" json:"hostUrl" yaml:"hostUrl"`
}

