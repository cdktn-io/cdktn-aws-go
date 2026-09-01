package awsecs


// Experimental.
type AwsEcsDaemonTaskDefinition_MountPointProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#container_path AwsEcsDaemonTaskDefinition#container_path}.
	// Experimental.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#read_only AwsEcsDaemonTaskDefinition#read_only}.
	// Experimental.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#source_volume AwsEcsDaemonTaskDefinition#source_volume}.
	// Experimental.
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

