package awssagemakerai


// Experimental.
type TfDomain_DefaultSpaceSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#execution_role TfDomain#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// custom_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#custom_file_system_config TfDomain#custom_file_system_config}
	// Experimental.
	CustomFileSystemConfig interface{} `field:"optional" json:"customFileSystemConfig" yaml:"customFileSystemConfig"`
	// custom_posix_user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#custom_posix_user_config TfDomain#custom_posix_user_config}
	// Experimental.
	CustomPosixUserConfig *TfDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// jupyter_lab_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#jupyter_lab_app_settings TfDomain#jupyter_lab_app_settings}
	// Experimental.
	JupyterLabAppSettings *TfDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#jupyter_server_app_settings TfDomain#jupyter_server_app_settings}
	// Experimental.
	JupyterServerAppSettings *TfDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#kernel_gateway_app_settings TfDomain#kernel_gateway_app_settings}
	// Experimental.
	KernelGatewayAppSettings *TfDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#security_groups TfDomain#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// space_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#space_storage_settings TfDomain#space_storage_settings}
	// Experimental.
	SpaceStorageSettings *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
}

