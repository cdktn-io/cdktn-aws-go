package awsbedrockagents


// Experimental.
type AwsBedrockagentDataSource_DataSourceConfigurationSharePointConfigurationSourceConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#auth_type AwsBedrockagentDataSource#auth_type}.
	// Experimental.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#credentials_secret_arn AwsBedrockagentDataSource#credentials_secret_arn}.
	// Experimental.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#domain AwsBedrockagentDataSource#domain}.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#host_type AwsBedrockagentDataSource#host_type}.
	// Experimental.
	HostType *string `field:"required" json:"hostType" yaml:"hostType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#site_urls AwsBedrockagentDataSource#site_urls}.
	// Experimental.
	SiteUrls *[]*string `field:"required" json:"siteUrls" yaml:"siteUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#tenant_id AwsBedrockagentDataSource#tenant_id}.
	// Experimental.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

