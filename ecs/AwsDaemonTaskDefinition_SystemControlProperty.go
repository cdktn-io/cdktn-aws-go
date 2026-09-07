package ecs


// Experimental.
type AwsDaemonTaskDefinition_SystemControlProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#namespace AwsDaemonTaskDefinition#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#value AwsDaemonTaskDefinition#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

