package sagemakerai


// Experimental.
type AwsUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#idle_timeout_in_minutes AwsUserProfile#idle_timeout_in_minutes}.
	// Experimental.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#lifecycle_management AwsUserProfile#lifecycle_management}.
	// Experimental.
	LifecycleManagement *string `field:"optional" json:"lifecycleManagement" yaml:"lifecycleManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#max_idle_timeout_in_minutes AwsUserProfile#max_idle_timeout_in_minutes}.
	// Experimental.
	MaxIdleTimeoutInMinutes *float64 `field:"optional" json:"maxIdleTimeoutInMinutes" yaml:"maxIdleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#min_idle_timeout_in_minutes AwsUserProfile#min_idle_timeout_in_minutes}.
	// Experimental.
	MinIdleTimeoutInMinutes *float64 `field:"optional" json:"minIdleTimeoutInMinutes" yaml:"minIdleTimeoutInMinutes"`
}

