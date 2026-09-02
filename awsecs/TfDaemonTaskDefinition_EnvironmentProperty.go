package awsecs


// Experimental.
type TfDaemonTaskDefinition_EnvironmentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#name TfDaemonTaskDefinition#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#value TfDaemonTaskDefinition#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

