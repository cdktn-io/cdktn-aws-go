package sagemakerai


// Experimental.
type AwsAppImageConfig_JupyterLabImageConfigProperty struct {
	// container_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_config AwsAppImageConfig#container_config}
	// Experimental.
	ContainerConfig *AwsAppImageConfig_JupyterLabImageConfigContainerConfigProperty `field:"optional" json:"containerConfig" yaml:"containerConfig"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#file_system_config AwsAppImageConfig#file_system_config}
	// Experimental.
	FileSystemConfig *AwsAppImageConfig_JupyterLabImageConfigFileSystemConfigProperty `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

