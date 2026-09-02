package awssagemakerai


// Experimental.
type TfUserProfile_IdentityProviderOauthSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#secret_arn TfUserProfile#secret_arn}.
	// Experimental.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#data_source_name TfUserProfile#data_source_name}.
	// Experimental.
	DataSourceName *string `field:"optional" json:"dataSourceName" yaml:"dataSourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#status TfUserProfile#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

