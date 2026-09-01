package awssagemakerai


// Experimental.
type AwsSagemakerUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#idle_timeout_in_minutes AwsSagemakerUserProfile#idle_timeout_in_minutes}.
	// Experimental.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#lifecycle_management AwsSagemakerUserProfile#lifecycle_management}.
	// Experimental.
	LifecycleManagement *string `field:"optional" json:"lifecycleManagement" yaml:"lifecycleManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#max_idle_timeout_in_minutes AwsSagemakerUserProfile#max_idle_timeout_in_minutes}.
	// Experimental.
	MaxIdleTimeoutInMinutes *float64 `field:"optional" json:"maxIdleTimeoutInMinutes" yaml:"maxIdleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#min_idle_timeout_in_minutes AwsSagemakerUserProfile#min_idle_timeout_in_minutes}.
	// Experimental.
	MinIdleTimeoutInMinutes *float64 `field:"optional" json:"minIdleTimeoutInMinutes" yaml:"minIdleTimeoutInMinutes"`
}

