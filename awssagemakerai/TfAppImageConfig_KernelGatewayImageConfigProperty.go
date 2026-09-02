package awssagemakerai


// Experimental.
type TfAppImageConfig_KernelGatewayImageConfigProperty struct {
	// kernel_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#kernel_spec TfAppImageConfig#kernel_spec}
	// Experimental.
	KernelSpec interface{} `field:"required" json:"kernelSpec" yaml:"kernelSpec"`
	// file_system_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#file_system_config TfAppImageConfig#file_system_config}
	// Experimental.
	FileSystemConfig *TfAppImageConfig_KernelGatewayImageConfigFileSystemConfigProperty `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

