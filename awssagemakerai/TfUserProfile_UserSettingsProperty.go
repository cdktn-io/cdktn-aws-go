package awssagemakerai


// Experimental.
type TfUserProfile_UserSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#execution_role TfUserProfile#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#auto_mount_home_efs TfUserProfile#auto_mount_home_efs}.
	// Experimental.
	AutoMountHomeEfs *string `field:"optional" json:"autoMountHomeEfs" yaml:"autoMountHomeEfs"`
	// canvas_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#canvas_app_settings TfUserProfile#canvas_app_settings}
	// Experimental.
	CanvasAppSettings *TfUserProfile_CanvasAppSettingsProperty `field:"optional" json:"canvasAppSettings" yaml:"canvasAppSettings"`
	// code_editor_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#code_editor_app_settings TfUserProfile#code_editor_app_settings}
	// Experimental.
	CodeEditorAppSettings *TfUserProfile_CodeEditorAppSettingsProperty `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// custom_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#custom_file_system_config TfUserProfile#custom_file_system_config}
	// Experimental.
	CustomFileSystemConfig interface{} `field:"optional" json:"customFileSystemConfig" yaml:"customFileSystemConfig"`
	// custom_posix_user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#custom_posix_user_config TfUserProfile#custom_posix_user_config}
	// Experimental.
	CustomPosixUserConfig *TfUserProfile_CustomPosixUserConfigProperty `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#default_landing_uri TfUserProfile#default_landing_uri}.
	// Experimental.
	DefaultLandingUri *string `field:"optional" json:"defaultLandingUri" yaml:"defaultLandingUri"`
	// jupyter_lab_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#jupyter_lab_app_settings TfUserProfile#jupyter_lab_app_settings}
	// Experimental.
	JupyterLabAppSettings *TfUserProfile_JupyterLabAppSettingsProperty `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#jupyter_server_app_settings TfUserProfile#jupyter_server_app_settings}
	// Experimental.
	JupyterServerAppSettings *TfUserProfile_JupyterServerAppSettingsProperty `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#kernel_gateway_app_settings TfUserProfile#kernel_gateway_app_settings}
	// Experimental.
	KernelGatewayAppSettings *TfUserProfile_KernelGatewayAppSettingsProperty `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// r_session_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#r_session_app_settings TfUserProfile#r_session_app_settings}
	// Experimental.
	RSessionAppSettings *TfUserProfile_RSessionAppSettingsProperty `field:"optional" json:"rSessionAppSettings" yaml:"rSessionAppSettings"`
	// r_studio_server_pro_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#r_studio_server_pro_app_settings TfUserProfile#r_studio_server_pro_app_settings}
	// Experimental.
	RStudioServerProAppSettings *TfUserProfile_RStudioServerProAppSettingsProperty `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#security_groups TfUserProfile#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// sharing_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#sharing_settings TfUserProfile#sharing_settings}
	// Experimental.
	SharingSettings *TfUserProfile_SharingSettingsProperty `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// space_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#space_storage_settings TfUserProfile#space_storage_settings}
	// Experimental.
	SpaceStorageSettings *TfUserProfile_SpaceStorageSettingsProperty `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#studio_web_portal TfUserProfile#studio_web_portal}.
	// Experimental.
	StudioWebPortal *string `field:"optional" json:"studioWebPortal" yaml:"studioWebPortal"`
	// studio_web_portal_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#studio_web_portal_settings TfUserProfile#studio_web_portal_settings}
	// Experimental.
	StudioWebPortalSettings *TfUserProfile_StudioWebPortalSettingsProperty `field:"optional" json:"studioWebPortalSettings" yaml:"studioWebPortalSettings"`
	// tensor_board_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#tensor_board_app_settings TfUserProfile#tensor_board_app_settings}
	// Experimental.
	TensorBoardAppSettings *TfUserProfile_TensorBoardAppSettingsProperty `field:"optional" json:"tensorBoardAppSettings" yaml:"tensorBoardAppSettings"`
}

