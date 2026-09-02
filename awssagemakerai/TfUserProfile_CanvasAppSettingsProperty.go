package awssagemakerai


// Experimental.
type TfUserProfile_CanvasAppSettingsProperty struct {
	// direct_deploy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#direct_deploy_settings TfUserProfile#direct_deploy_settings}
	// Experimental.
	DirectDeploySettings *TfUserProfile_DirectDeploySettingsProperty `field:"optional" json:"directDeploySettings" yaml:"directDeploySettings"`
	// emr_serverless_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#emr_serverless_settings TfUserProfile#emr_serverless_settings}
	// Experimental.
	EmrServerlessSettings *TfUserProfile_EmrServerlessSettingsProperty `field:"optional" json:"emrServerlessSettings" yaml:"emrServerlessSettings"`
	// generative_ai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#generative_ai_settings TfUserProfile#generative_ai_settings}
	// Experimental.
	GenerativeAiSettings *TfUserProfile_GenerativeAiSettingsProperty `field:"optional" json:"generativeAiSettings" yaml:"generativeAiSettings"`
	// identity_provider_oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#identity_provider_oauth_settings TfUserProfile#identity_provider_oauth_settings}
	// Experimental.
	IdentityProviderOauthSettings interface{} `field:"optional" json:"identityProviderOauthSettings" yaml:"identityProviderOauthSettings"`
	// kendra_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#kendra_settings TfUserProfile#kendra_settings}
	// Experimental.
	KendraSettings *TfUserProfile_KendraSettingsProperty `field:"optional" json:"kendraSettings" yaml:"kendraSettings"`
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#model_register_settings TfUserProfile#model_register_settings}
	// Experimental.
	ModelRegisterSettings *TfUserProfile_ModelRegisterSettingsProperty `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#time_series_forecasting_settings TfUserProfile#time_series_forecasting_settings}
	// Experimental.
	TimeSeriesForecastingSettings *TfUserProfile_TimeSeriesForecastingSettingsProperty `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
	// workspace_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#workspace_settings TfUserProfile#workspace_settings}
	// Experimental.
	WorkspaceSettings *TfUserProfile_WorkspaceSettingsProperty `field:"optional" json:"workspaceSettings" yaml:"workspaceSettings"`
}

