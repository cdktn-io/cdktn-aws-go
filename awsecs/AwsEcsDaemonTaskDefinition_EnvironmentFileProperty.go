package awsecs


// Experimental.
type AwsEcsDaemonTaskDefinition_EnvironmentFileProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#type AwsEcsDaemonTaskDefinition#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#value AwsEcsDaemonTaskDefinition#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

