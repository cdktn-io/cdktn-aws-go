package sagemakerai


// Experimental.
type AwsDomain_DefaultSpaceSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#execution_role AwsDomain#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// custom_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#custom_file_system_config AwsDomain#custom_file_system_config}
	// Experimental.
	CustomFileSystemConfig interface{} `field:"optional" json:"customFileSystemConfig" yaml:"customFileSystemConfig"`
	// custom_posix_user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#custom_posix_user_config AwsDomain#custom_posix_user_config}
	// Experimental.
	CustomPosixUserConfig *AwsDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// jupyter_lab_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#jupyter_lab_app_settings AwsDomain#jupyter_lab_app_settings}
	// Experimental.
	JupyterLabAppSettings *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#jupyter_server_app_settings AwsDomain#jupyter_server_app_settings}
	// Experimental.
	JupyterServerAppSettings *AwsDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#kernel_gateway_app_settings AwsDomain#kernel_gateway_app_settings}
	// Experimental.
	KernelGatewayAppSettings *AwsDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#security_groups AwsDomain#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// space_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#space_storage_settings AwsDomain#space_storage_settings}
	// Experimental.
	SpaceStorageSettings *AwsDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
}

