package awssagemakerai


// Experimental.
type TfUserProfile_CodeEditorAppSettingsProperty struct {
	// app_lifecycle_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#app_lifecycle_management TfUserProfile#app_lifecycle_management}
	// Experimental.
	AppLifecycleManagement *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#built_in_lifecycle_config_arn TfUserProfile#built_in_lifecycle_config_arn}.
	// Experimental.
	BuiltInLifecycleConfigArn *string `field:"optional" json:"builtInLifecycleConfigArn" yaml:"builtInLifecycleConfigArn"`
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#custom_image TfUserProfile#custom_image}
	// Experimental.
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#default_resource_spec TfUserProfile#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *TfUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#lifecycle_config_arns TfUserProfile#lifecycle_config_arns}.
	// Experimental.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

