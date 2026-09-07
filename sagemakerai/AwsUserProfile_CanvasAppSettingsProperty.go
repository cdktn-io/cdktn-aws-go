package sagemakerai


// Experimental.
type AwsUserProfile_CanvasAppSettingsProperty struct {
	// direct_deploy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#direct_deploy_settings AwsUserProfile#direct_deploy_settings}
	// Experimental.
	DirectDeploySettings *AwsUserProfile_DirectDeploySettingsProperty `field:"optional" json:"directDeploySettings" yaml:"directDeploySettings"`
	// emr_serverless_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#emr_serverless_settings AwsUserProfile#emr_serverless_settings}
	// Experimental.
	EmrServerlessSettings *AwsUserProfile_EmrServerlessSettingsProperty `field:"optional" json:"emrServerlessSettings" yaml:"emrServerlessSettings"`
	// generative_ai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#generative_ai_settings AwsUserProfile#generative_ai_settings}
	// Experimental.
	GenerativeAiSettings *AwsUserProfile_GenerativeAiSettingsProperty `field:"optional" json:"generativeAiSettings" yaml:"generativeAiSettings"`
	// identity_provider_oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#identity_provider_oauth_settings AwsUserProfile#identity_provider_oauth_settings}
	// Experimental.
	IdentityProviderOauthSettings interface{} `field:"optional" json:"identityProviderOauthSettings" yaml:"identityProviderOauthSettings"`
	// kendra_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#kendra_settings AwsUserProfile#kendra_settings}
	// Experimental.
	KendraSettings *AwsUserProfile_KendraSettingsProperty `field:"optional" json:"kendraSettings" yaml:"kendraSettings"`
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#model_register_settings AwsUserProfile#model_register_settings}
	// Experimental.
	ModelRegisterSettings *AwsUserProfile_ModelRegisterSettingsProperty `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#time_series_forecasting_settings AwsUserProfile#time_series_forecasting_settings}
	// Experimental.
	TimeSeriesForecastingSettings *AwsUserProfile_TimeSeriesForecastingSettingsProperty `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
	// workspace_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#workspace_settings AwsUserProfile#workspace_settings}
	// Experimental.
	WorkspaceSettings *AwsUserProfile_WorkspaceSettingsProperty `field:"optional" json:"workspaceSettings" yaml:"workspaceSettings"`
}

