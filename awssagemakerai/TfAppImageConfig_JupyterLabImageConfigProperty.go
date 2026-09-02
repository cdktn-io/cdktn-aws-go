package awssagemakerai


// Experimental.
type TfAppImageConfig_JupyterLabImageConfigProperty struct {
	// container_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_config TfAppImageConfig#container_config}
	// Experimental.
	ContainerConfig *TfAppImageConfig_JupyterLabImageConfigContainerConfigProperty `field:"optional" json:"containerConfig" yaml:"containerConfig"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#file_system_config TfAppImageConfig#file_system_config}
	// Experimental.
	FileSystemConfig *TfAppImageConfig_JupyterLabImageConfigFileSystemConfigProperty `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

