package awssagemakerai


// Experimental.
type TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_arguments TfAppImageConfig#container_arguments}.
	// Experimental.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_entrypoint TfAppImageConfig#container_entrypoint}.
	// Experimental.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_environment_variables TfAppImageConfig#container_environment_variables}.
	// Experimental.
	ContainerEnvironmentVariables *map[string]*string `field:"optional" json:"containerEnvironmentVariables" yaml:"containerEnvironmentVariables"`
}

