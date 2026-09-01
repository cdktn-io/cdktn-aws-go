package awssagemakerai


// Experimental.
type AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty struct {
	// idle_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#idle_settings AwsSagemakerUserProfile#idle_settings}
	// Experimental.
	IdleSettings *AwsSagemakerUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty `field:"optional" json:"idleSettings" yaml:"idleSettings"`
}

