package awssagemakerai


// Experimental.
type AwsSagemakerAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_arguments AwsSagemakerAppImageConfig#container_arguments}.
	// Experimental.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_entrypoint AwsSagemakerAppImageConfig#container_entrypoint}.
	// Experimental.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_environment_variables AwsSagemakerAppImageConfig#container_environment_variables}.
	// Experimental.
	ContainerEnvironmentVariables *map[string]*string `field:"optional" json:"containerEnvironmentVariables" yaml:"containerEnvironmentVariables"`
}

