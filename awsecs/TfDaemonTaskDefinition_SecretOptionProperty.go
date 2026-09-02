package awsecs


// Experimental.
type TfDaemonTaskDefinition_SecretOptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#name TfDaemonTaskDefinition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#value_from TfDaemonTaskDefinition#value_from}.
	// Experimental.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

