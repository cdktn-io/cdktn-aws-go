package ecs


// Experimental.
type AwsDaemonTaskDefinition_SecretProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#name AwsDaemonTaskDefinition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#value_from AwsDaemonTaskDefinition#value_from}.
	// Experimental.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

