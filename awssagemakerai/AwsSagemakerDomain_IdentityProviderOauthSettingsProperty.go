package awssagemakerai


// Experimental.
type AwsSagemakerDomain_IdentityProviderOauthSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#secret_arn AwsSagemakerDomain#secret_arn}.
	// Experimental.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#data_source_name AwsSagemakerDomain#data_source_name}.
	// Experimental.
	DataSourceName *string `field:"optional" json:"dataSourceName" yaml:"dataSourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#status AwsSagemakerDomain#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

