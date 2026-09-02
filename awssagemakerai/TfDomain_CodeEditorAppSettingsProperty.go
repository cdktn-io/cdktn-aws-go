package awssagemakerai


// Experimental.
type TfDomain_CodeEditorAppSettingsProperty struct {
	// app_lifecycle_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#app_lifecycle_management TfDomain#app_lifecycle_management}
	// Experimental.
	AppLifecycleManagement *TfDomain_DefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#built_in_lifecycle_config_arn TfDomain#built_in_lifecycle_config_arn}.
	// Experimental.
	BuiltInLifecycleConfigArn *string `field:"optional" json:"builtInLifecycleConfigArn" yaml:"builtInLifecycleConfigArn"`
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#custom_image TfDomain#custom_image}
	// Experimental.
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#default_resource_spec TfDomain#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *TfDomain_DefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#lifecycle_config_arns TfDomain#lifecycle_config_arns}.
	// Experimental.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

