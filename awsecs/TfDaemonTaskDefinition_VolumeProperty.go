package awsecs


// Experimental.
type TfDaemonTaskDefinition_VolumeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#name TfDaemonTaskDefinition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#host TfDaemonTaskDefinition#host}
	// Experimental.
	Host interface{} `field:"optional" json:"host" yaml:"host"`
}

