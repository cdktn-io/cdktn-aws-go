package awsecs


// Experimental.
type TfDaemonTaskDefinition_SystemControlProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#namespace TfDaemonTaskDefinition#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#value TfDaemonTaskDefinition#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

