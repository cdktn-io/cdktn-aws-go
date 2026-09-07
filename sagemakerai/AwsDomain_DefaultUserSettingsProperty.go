package sagemakerai


// Experimental.
type AwsDomain_DefaultUserSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#execution_role AwsDomain#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#auto_mount_home_efs AwsDomain#auto_mount_home_efs}.
	// Experimental.
	AutoMountHomeEfs *string `field:"optional" json:"autoMountHomeEfs" yaml:"autoMountHomeEfs"`
	// canvas_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#canvas_app_settings AwsDomain#canvas_app_settings}
	// Experimental.
	CanvasAppSettings *AwsDomain_CanvasAppSettingsProperty `field:"optional" json:"canvasAppSettings" yaml:"canvasAppSettings"`
	// code_editor_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#code_editor_app_settings AwsDomain#code_editor_app_settings}
	// Experimental.
	CodeEditorAppSettings *AwsDomain_CodeEditorAppSettingsProperty `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// custom_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#custom_file_system_config AwsDomain#custom_file_system_config}
	// Experimental.
	CustomFileSystemConfig interface{} `field:"optional" json:"customFileSystemConfig" yaml:"customFileSystemConfig"`
	// custom_posix_user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#custom_posix_user_config AwsDomain#custom_posix_user_config}
	// Experimental.
	CustomPosixUserConfig *AwsDomain_DefaultUserSettingsCustomPosixUserConfigProperty `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#default_landing_uri AwsDomain#default_landing_uri}.
	// Experimental.
	DefaultLandingUri *string `field:"optional" json:"defaultLandingUri" yaml:"defaultLandingUri"`
	// jupyter_lab_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#jupyter_lab_app_settings AwsDomain#jupyter_lab_app_settings}
	// Experimental.
	JupyterLabAppSettings *AwsDomain_DefaultUserSettingsJupyterLabAppSettingsProperty `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#jupyter_server_app_settings AwsDomain#jupyter_server_app_settings}
	// Experimental.
	JupyterServerAppSettings *AwsDomain_DefaultUserSettingsJupyterServerAppSettingsProperty `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#kernel_gateway_app_settings AwsDomain#kernel_gateway_app_settings}
	// Experimental.
	KernelGatewayAppSettings *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// r_session_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#r_session_app_settings AwsDomain#r_session_app_settings}
	// Experimental.
	RSessionAppSettings *AwsDomain_RSessionAppSettingsProperty `field:"optional" json:"rSessionAppSettings" yaml:"rSessionAppSettings"`
	// r_studio_server_pro_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#r_studio_server_pro_app_settings AwsDomain#r_studio_server_pro_app_settings}
	// Experimental.
	RStudioServerProAppSettings *AwsDomain_RStudioServerProAppSettingsProperty `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#security_groups AwsDomain#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// sharing_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#sharing_settings AwsDomain#sharing_settings}
	// Experimental.
	SharingSettings *AwsDomain_SharingSettingsProperty `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// space_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#space_storage_settings AwsDomain#space_storage_settings}
	// Experimental.
	SpaceStorageSettings *AwsDomain_DefaultUserSettingsSpaceStorageSettingsProperty `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#studio_web_portal AwsDomain#studio_web_portal}.
	// Experimental.
	StudioWebPortal *string `field:"optional" json:"studioWebPortal" yaml:"studioWebPortal"`
	// studio_web_portal_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#studio_web_portal_settings AwsDomain#studio_web_portal_settings}
	// Experimental.
	StudioWebPortalSettings *AwsDomain_StudioWebPortalSettingsProperty `field:"optional" json:"studioWebPortalSettings" yaml:"studioWebPortalSettings"`
	// tensor_board_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#tensor_board_app_settings AwsDomain#tensor_board_app_settings}
	// Experimental.
	TensorBoardAppSettings *AwsDomain_TensorBoardAppSettingsProperty `field:"optional" json:"tensorBoardAppSettings" yaml:"tensorBoardAppSettings"`
}

