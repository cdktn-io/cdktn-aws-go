package awssagemakerai


// Experimental.
type AwsSagemakerDomain_CanvasAppSettingsProperty struct {
	// direct_deploy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#direct_deploy_settings AwsSagemakerDomain#direct_deploy_settings}
	// Experimental.
	DirectDeploySettings *AwsSagemakerDomain_DirectDeploySettingsProperty `field:"optional" json:"directDeploySettings" yaml:"directDeploySettings"`
	// emr_serverless_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#emr_serverless_settings AwsSagemakerDomain#emr_serverless_settings}
	// Experimental.
	EmrServerlessSettings *AwsSagemakerDomain_EmrServerlessSettingsProperty `field:"optional" json:"emrServerlessSettings" yaml:"emrServerlessSettings"`
	// generative_ai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#generative_ai_settings AwsSagemakerDomain#generative_ai_settings}
	// Experimental.
	GenerativeAiSettings *AwsSagemakerDomain_GenerativeAiSettingsProperty `field:"optional" json:"generativeAiSettings" yaml:"generativeAiSettings"`
	// identity_provider_oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#identity_provider_oauth_settings AwsSagemakerDomain#identity_provider_oauth_settings}
	// Experimental.
	IdentityProviderOauthSettings interface{} `field:"optional" json:"identityProviderOauthSettings" yaml:"identityProviderOauthSettings"`
	// kendra_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#kendra_settings AwsSagemakerDomain#kendra_settings}
	// Experimental.
	KendraSettings *AwsSagemakerDomain_KendraSettingsProperty `field:"optional" json:"kendraSettings" yaml:"kendraSettings"`
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#model_register_settings AwsSagemakerDomain#model_register_settings}
	// Experimental.
	ModelRegisterSettings *AwsSagemakerDomain_ModelRegisterSettingsProperty `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#time_series_forecasting_settings AwsSagemakerDomain#time_series_forecasting_settings}
	// Experimental.
	TimeSeriesForecastingSettings *AwsSagemakerDomain_TimeSeriesForecastingSettingsProperty `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
	// workspace_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#workspace_settings AwsSagemakerDomain#workspace_settings}
	// Experimental.
	WorkspaceSettings *AwsSagemakerDomain_WorkspaceSettingsProperty `field:"optional" json:"workspaceSettings" yaml:"workspaceSettings"`
}

