package awsecs


// Experimental.
type TfTaskDefinition_EfsVolumeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#file_system_id TfTaskDefinition#file_system_id}.
	// Experimental.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#authorization_config TfTaskDefinition#authorization_config}
	// Experimental.
	AuthorizationConfig *TfTaskDefinition_VolumeEfsVolumeConfigurationAuthorizationConfigProperty `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#root_directory TfTaskDefinition#root_directory}.
	// Experimental.
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#transit_encryption TfTaskDefinition#transit_encryption}.
	// Experimental.
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#transit_encryption_port TfTaskDefinition#transit_encryption_port}.
	// Experimental.
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

