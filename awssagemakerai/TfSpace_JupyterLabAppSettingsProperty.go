package awssagemakerai


// Experimental.
type TfSpace_JupyterLabAppSettingsProperty struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#default_resource_spec TfSpace#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *TfSpace_SpaceSettingsJupyterLabAppSettingsDefaultResourceSpecProperty `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// app_lifecycle_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#app_lifecycle_management TfSpace#app_lifecycle_management}
	// Experimental.
	AppLifecycleManagement *TfSpace_SpaceSettingsJupyterLabAppSettingsAppLifecycleManagementProperty `field:"optional" json:"appLifecycleManagement" yaml:"appLifecycleManagement"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#code_repository TfSpace#code_repository}
	// Experimental.
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
}

