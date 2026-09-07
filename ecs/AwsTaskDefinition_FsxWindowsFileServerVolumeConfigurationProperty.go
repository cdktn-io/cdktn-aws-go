package ecs


// Experimental.
type AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty struct {
	// authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#authorization_config AwsTaskDefinition#authorization_config}
	// Experimental.
	AuthorizationConfig *AwsTaskDefinition_VolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigProperty `field:"required" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#file_system_id AwsTaskDefinition#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#root_directory AwsTaskDefinition#root_directory}.
	// Experimental.
	RootDirectory *string `field:"required" json:"rootDirectory" yaml:"rootDirectory"`
}

