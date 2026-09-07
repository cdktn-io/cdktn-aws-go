package sagemakerai


// Experimental.
type AwsUserProfile_UserSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#execution_role AwsUserProfile#execution_role}.
	// Experimental.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#auto_mount_home_efs AwsUserProfile#auto_mount_home_efs}.
	// Experimental.
	AutoMountHomeEfs *string `field:"optional" json:"autoMountHomeEfs" yaml:"autoMountHomeEfs"`
	// canvas_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#canvas_app_settings AwsUserProfile#canvas_app_settings}
	// Experimental.
	CanvasAppSettings *AwsUserProfile_CanvasAppSettingsProperty `field:"optional" json:"canvasAppSettings" yaml:"canvasAppSettings"`
	// code_editor_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#code_editor_app_settings AwsUserProfile#code_editor_app_settings}
	// Experimental.
	CodeEditorAppSettings *AwsUserProfile_CodeEditorAppSettingsProperty `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// custom_file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#custom_file_system_config AwsUserProfile#custom_file_system_config}
	// Experimental.
	CustomFileSystemConfig interface{} `field:"optional" json:"customFileSystemConfig" yaml:"customFileSystemConfig"`
	// custom_posix_user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#custom_posix_user_config AwsUserProfile#custom_posix_user_config}
	// Experimental.
	CustomPosixUserConfig *AwsUserProfile_CustomPosixUserConfigProperty `field:"optional" json:"customPosixUserConfig" yaml:"customPosixUserConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#default_landing_uri AwsUserProfile#default_landing_uri}.
	// Experimental.
	DefaultLandingUri *string `field:"optional" json:"defaultLandingUri" yaml:"defaultLandingUri"`
	// jupyter_lab_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#jupyter_lab_app_settings AwsUserProfile#jupyter_lab_app_settings}
	// Experimental.
	JupyterLabAppSettings *AwsUserProfile_JupyterLabAppSettingsProperty `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// jupyter_server_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#jupyter_server_app_settings AwsUserProfile#jupyter_server_app_settings}
	// Experimental.
	JupyterServerAppSettings *AwsUserProfile_JupyterServerAppSettingsProperty `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// kernel_gateway_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#kernel_gateway_app_settings AwsUserProfile#kernel_gateway_app_settings}
	// Experimental.
	KernelGatewayAppSettings *AwsUserProfile_KernelGatewayAppSettingsProperty `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// r_session_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#r_session_app_settings AwsUserProfile#r_session_app_settings}
	// Experimental.
	RSessionAppSettings *AwsUserProfile_RSessionAppSettingsProperty `field:"optional" json:"rSessionAppSettings" yaml:"rSessionAppSettings"`
	// r_studio_server_pro_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#r_studio_server_pro_app_settings AwsUserProfile#r_studio_server_pro_app_settings}
	// Experimental.
	RStudioServerProAppSettings *AwsUserProfile_RStudioServerProAppSettingsProperty `field:"optional" json:"rStudioServerProAppSettings" yaml:"rStudioServerProAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#security_groups AwsUserProfile#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// sharing_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#sharing_settings AwsUserProfile#sharing_settings}
	// Experimental.
	SharingSettings *AwsUserProfile_SharingSettingsProperty `field:"optional" json:"sharingSettings" yaml:"sharingSettings"`
	// space_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#space_storage_settings AwsUserProfile#space_storage_settings}
	// Experimental.
	SpaceStorageSettings *AwsUserProfile_SpaceStorageSettingsProperty `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#studio_web_portal AwsUserProfile#studio_web_portal}.
	// Experimental.
	StudioWebPortal *string `field:"optional" json:"studioWebPortal" yaml:"studioWebPortal"`
	// studio_web_portal_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#studio_web_portal_settings AwsUserProfile#studio_web_portal_settings}
	// Experimental.
	StudioWebPortalSettings *AwsUserProfile_StudioWebPortalSettingsProperty `field:"optional" json:"studioWebPortalSettings" yaml:"studioWebPortalSettings"`
	// tensor_board_app_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#tensor_board_app_settings AwsUserProfile#tensor_board_app_settings}
	// Experimental.
	TensorBoardAppSettings *AwsUserProfile_TensorBoardAppSettingsProperty `field:"optional" json:"tensorBoardAppSettings" yaml:"tensorBoardAppSettings"`
}

