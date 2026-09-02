package awsecs


// Experimental.
type TfDaemonTaskDefinition_FirelensConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#type TfDaemonTaskDefinition#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#options TfDaemonTaskDefinition#options}.
	// Experimental.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
}

