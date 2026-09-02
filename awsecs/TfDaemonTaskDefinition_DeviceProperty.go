package awsecs


// Experimental.
type TfDaemonTaskDefinition_DeviceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#host_path TfDaemonTaskDefinition#host_path}.
	// Experimental.
	HostPath *string `field:"required" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#container_path TfDaemonTaskDefinition#container_path}.
	// Experimental.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#permissions TfDaemonTaskDefinition#permissions}.
	// Experimental.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

