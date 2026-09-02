package awssagemakerai


// Experimental.
type TfSpace_CodeEditorAppSettingsProperty struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#default_resource_spec TfSpace#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *TfSpace_SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecProperty `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// app_lifecycle_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#app_lifecycle_management TfSpace#app_lifecycle_management}
	// Experimental.
	AppLifecycleManagement *TfSpace_SpaceSettingsCodeEditorAppSettingsAppLifecycleManagementProperty `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
}

