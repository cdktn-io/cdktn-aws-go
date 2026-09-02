package awsecs


// Experimental.
type TfTaskDefinition_S3filesVolumeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#file_system_arn TfTaskDefinition#file_system_arn}.
	// Experimental.
	FileSystemArn *string `field:"required" json:"fileSystemArn" yaml:"fileSystemArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#access_point_arn TfTaskDefinition#access_point_arn}.
	// Experimental.
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#root_directory TfTaskDefinition#root_directory}.
	// Experimental.
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#transit_encryption_port TfTaskDefinition#transit_encryption_port}.
	// Experimental.
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

