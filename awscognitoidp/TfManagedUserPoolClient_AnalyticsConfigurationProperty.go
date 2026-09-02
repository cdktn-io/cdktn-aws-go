package awscognitoidp


// Experimental.
type TfManagedUserPoolClient_AnalyticsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#application_arn TfManagedUserPoolClient#application_arn}.
	// Experimental.
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#application_id TfManagedUserPoolClient#application_id}.
	// Experimental.
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#external_id TfManagedUserPoolClient#external_id}.
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#role_arn TfManagedUserPoolClient#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_managed_user_pool_client#user_data_shared TfManagedUserPoolClient#user_data_shared}.
	// Experimental.
	UserDataShared interface{} `field:"optional" json:"userDataShared" yaml:"userDataShared"`
}

