package awssagemakerai


// Experimental.
type AwsSagemakerAppImageConfig_CodeEditorAppImageConfigProperty struct {
	// container_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_config AwsSagemakerAppImageConfig#container_config}
	// Experimental.
	ContainerConfig *AwsSagemakerAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty `field:"optional" json:"containerConfig" yaml:"containerConfig"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#file_system_config AwsSagemakerAppImageConfig#file_system_config}
	// Experimental.
	FileSystemConfig *AwsSagemakerAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

