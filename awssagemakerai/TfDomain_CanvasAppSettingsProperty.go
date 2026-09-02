package awssagemakerai


// Experimental.
type TfDomain_CanvasAppSettingsProperty struct {
	// direct_deploy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#direct_deploy_settings TfDomain#direct_deploy_settings}
	// Experimental.
	DirectDeploySettings *TfDomain_DirectDeploySettingsProperty `field:"optional" json:"directDeploySettings" yaml:"directDeploySettings"`
	// emr_serverless_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#emr_serverless_settings TfDomain#emr_serverless_settings}
	// Experimental.
	EmrServerlessSettings *TfDomain_EmrServerlessSettingsProperty `field:"optional" json:"emrServerlessSettings" yaml:"emrServerlessSettings"`
	// generative_ai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#generative_ai_settings TfDomain#generative_ai_settings}
	// Experimental.
	GenerativeAiSettings *TfDomain_GenerativeAiSettingsProperty `field:"optional" json:"generativeAiSettings" yaml:"generativeAiSettings"`
	// identity_provider_oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#identity_provider_oauth_settings TfDomain#identity_provider_oauth_settings}
	// Experimental.
	IdentityProviderOauthSettings interface{} `field:"optional" json:"identityProviderOauthSettings" yaml:"identityProviderOauthSettings"`
	// kendra_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#kendra_settings TfDomain#kendra_settings}
	// Experimental.
	KendraSettings *TfDomain_KendraSettingsProperty `field:"optional" json:"kendraSettings" yaml:"kendraSettings"`
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#model_register_settings TfDomain#model_register_settings}
	// Experimental.
	ModelRegisterSettings *TfDomain_ModelRegisterSettingsProperty `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#time_series_forecasting_settings TfDomain#time_series_forecasting_settings}
	// Experimental.
	TimeSeriesForecastingSettings *TfDomain_TimeSeriesForecastingSettingsProperty `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
	// workspace_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#workspace_settings TfDomain#workspace_settings}
	// Experimental.
	WorkspaceSettings *TfDomain_WorkspaceSettingsProperty `field:"optional" json:"workspaceSettings" yaml:"workspaceSettings"`
}

