package ecs


// Experimental.
type AwsDaemonTaskDefinition_TmpfsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#container_path AwsDaemonTaskDefinition#container_path}.
	// Experimental.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#size AwsDaemonTaskDefinition#size}.
	// Experimental.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#mount_options AwsDaemonTaskDefinition#mount_options}.
	// Experimental.
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

