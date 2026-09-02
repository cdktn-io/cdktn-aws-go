package awssagemakerai


// Experimental.
type TfAppImageConfig_CodeEditorAppImageConfigProperty struct {
	// container_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_config TfAppImageConfig#container_config}
	// Experimental.
	ContainerConfig *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty `field:"optional" json:"containerConfig" yaml:"containerConfig"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#file_system_config TfAppImageConfig#file_system_config}
	// Experimental.
	FileSystemConfig *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

