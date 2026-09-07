package sagemakerai


// Experimental.
type AwsAppImageConfig_JupyterLabImageConfigContainerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_arguments AwsAppImageConfig#container_arguments}.
	// Experimental.
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_entrypoint AwsAppImageConfig#container_entrypoint}.
	// Experimental.
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_app_image_config#container_environment_variables AwsAppImageConfig#container_environment_variables}.
	// Experimental.
	ContainerEnvironmentVariables *map[string]*string `field:"optional" json:"containerEnvironmentVariables" yaml:"containerEnvironmentVariables"`
}

