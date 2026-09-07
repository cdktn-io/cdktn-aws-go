package sagemakerai


// Experimental.
type AwsSpace_CodeEditorAppSettingsProperty struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#default_resource_spec AwsSpace#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *AwsSpace_SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecProperty `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// app_lifecycle_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#app_lifecycle_management AwsSpace#app_lifecycle_management}
	// Experimental.
	AppLifecycleManagement *AwsSpace_SpaceSettingsCodeEditorAppSettingsAppLifecycleManagementProperty `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
}

