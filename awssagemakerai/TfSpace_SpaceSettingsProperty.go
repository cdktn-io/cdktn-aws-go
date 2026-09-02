package awssagemakerai


// Experimental.
type TfSpace_SpaceSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#app_type TfSpace#app_type}.
	// Experimental.
	AppType *string `field:"optional" json:"appType" yaml:"appType"`
	// code_editor_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#code_editor_app_settings TfSpace#code_editor_app_settings}
	// Experimental.
	CodeEditorAppSettings *TfSpace_CodeEditorAppSettingsProperty `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// custom_file_system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#custom_file_system TfSpace#custom_file_system}
	// Experimental.
	CustomFileSystem interface{} `field:"optional" json:"customFileSystem" yaml:"customFileSystem"`
	// jupyter_lab_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#jupyter_lab_app_settings TfSpace#jupyter_lab_app_settings}
	// Experimental.
	JupyterLabAppSettings *TfSpace_JupyterLabAppSettingsProperty `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#jupyter_server_app_settings TfSpace#jupyter_server_app_settings}
	// Experimental.
	JupyterServerAppSettings *TfSpace_JupyterServerAppSettingsProperty `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#kernel_gateway_app_settings TfSpace#kernel_gateway_app_settings}
	// Experimental.
	KernelGatewayAppSettings *TfSpace_KernelGatewayAppSettingsProperty `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// space_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#space_storage_settings TfSpace#space_storage_settings}
	// Experimental.
	SpaceStorageSettings *TfSpace_SpaceStorageSettingsProperty `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
}

