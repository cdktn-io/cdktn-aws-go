package awssagemakerai


// Experimental.
type AwsSagemakerSpace_CodeEditorAppSettingsProperty struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#default_resource_spec AwsSagemakerSpace#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *AwsSagemakerSpace_SpaceSettingsCodeEditorAppSettingsDefaultResourceSpecProperty `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// app_lifecycle_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#app_lifecycle_management AwsSagemakerSpace#app_lifecycle_management}
	// Experimental.
	AppLifecycleManagement *AwsSagemakerSpace_SpaceSettingsCodeEditorAppSettingsAppLifecycleManagementProperty `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
}

