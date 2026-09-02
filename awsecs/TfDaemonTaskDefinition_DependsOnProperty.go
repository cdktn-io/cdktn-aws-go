package awsecs


// Experimental.
type TfDaemonTaskDefinition_DependsOnProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#condition TfDaemonTaskDefinition#condition}.
	// Experimental.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#container_name TfDaemonTaskDefinition#container_name}.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
}

