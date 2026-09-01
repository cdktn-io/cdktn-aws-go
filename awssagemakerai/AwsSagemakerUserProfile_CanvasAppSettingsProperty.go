package awssagemakerai


// Experimental.
type AwsSagemakerUserProfile_CanvasAppSettingsProperty struct {
	// direct_deploy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#direct_deploy_settings AwsSagemakerUserProfile#direct_deploy_settings}
	// Experimental.
	DirectDeploySettings *AwsSagemakerUserProfile_DirectDeploySettingsProperty `field:"optional" json:"directDeploySettings" yaml:"directDeploySettings"`
	// emr_serverless_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#emr_serverless_settings AwsSagemakerUserProfile#emr_serverless_settings}
	// Experimental.
	EmrServerlessSettings *AwsSagemakerUserProfile_EmrServerlessSettingsProperty `field:"optional" json:"emrServerlessSettings" yaml:"emrServerlessSettings"`
	// generative_ai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#generative_ai_settings AwsSagemakerUserProfile#generative_ai_settings}
	// Experimental.
	GenerativeAiSettings *AwsSagemakerUserProfile_GenerativeAiSettingsProperty `field:"optional" json:"generativeAiSettings" yaml:"generativeAiSettings"`
	// identity_provider_oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#identity_provider_oauth_settings AwsSagemakerUserProfile#identity_provider_oauth_settings}
	// Experimental.
	IdentityProviderOauthSettings interface{} `field:"optional" json:"identityProviderOauthSettings" yaml:"identityProviderOauthSettings"`
	// kendra_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#kendra_settings AwsSagemakerUserProfile#kendra_settings}
	// Experimental.
	KendraSettings *AwsSagemakerUserProfile_KendraSettingsProperty `field:"optional" json:"kendraSettings" yaml:"kendraSettings"`
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#model_register_settings AwsSagemakerUserProfile#model_register_settings}
	// Experimental.
	ModelRegisterSettings *AwsSagemakerUserProfile_ModelRegisterSettingsProperty `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#time_series_forecasting_settings AwsSagemakerUserProfile#time_series_forecasting_settings}
	// Experimental.
	TimeSeriesForecastingSettings *AwsSagemakerUserProfile_TimeSeriesForecastingSettingsProperty `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
	// workspace_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#workspace_settings AwsSagemakerUserProfile#workspace_settings}
	// Experimental.
	WorkspaceSettings *AwsSagemakerUserProfile_WorkspaceSettingsProperty `field:"optional" json:"workspaceSettings" yaml:"workspaceSettings"`
}

