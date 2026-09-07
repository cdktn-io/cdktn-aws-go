package sagemakerai


// Experimental.
type AwsDomain_CanvasAppSettingsProperty struct {
	// direct_deploy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#direct_deploy_settings AwsDomain#direct_deploy_settings}
	// Experimental.
	DirectDeploySettings *AwsDomain_DirectDeploySettingsProperty `field:"optional" json:"directDeploySettings" yaml:"directDeploySettings"`
	// emr_serverless_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#emr_serverless_settings AwsDomain#emr_serverless_settings}
	// Experimental.
	EmrServerlessSettings *AwsDomain_EmrServerlessSettingsProperty `field:"optional" json:"emrServerlessSettings" yaml:"emrServerlessSettings"`
	// generative_ai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#generative_ai_settings AwsDomain#generative_ai_settings}
	// Experimental.
	GenerativeAiSettings *AwsDomain_GenerativeAiSettingsProperty `field:"optional" json:"generativeAiSettings" yaml:"generativeAiSettings"`
	// identity_provider_oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#identity_provider_oauth_settings AwsDomain#identity_provider_oauth_settings}
	// Experimental.
	IdentityProviderOauthSettings interface{} `field:"optional" json:"identityProviderOauthSettings" yaml:"identityProviderOauthSettings"`
	// kendra_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#kendra_settings AwsDomain#kendra_settings}
	// Experimental.
	KendraSettings *AwsDomain_KendraSettingsProperty `field:"optional" json:"kendraSettings" yaml:"kendraSettings"`
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#model_register_settings AwsDomain#model_register_settings}
	// Experimental.
	ModelRegisterSettings *AwsDomain_ModelRegisterSettingsProperty `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#time_series_forecasting_settings AwsDomain#time_series_forecasting_settings}
	// Experimental.
	TimeSeriesForecastingSettings *AwsDomain_TimeSeriesForecastingSettingsProperty `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
	// workspace_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#workspace_settings AwsDomain#workspace_settings}
	// Experimental.
	WorkspaceSettings *AwsDomain_WorkspaceSettingsProperty `field:"optional" json:"workspaceSettings" yaml:"workspaceSettings"`
}

