package ecs


// Experimental.
type AwsDaemonTaskDefinition_LogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#log_driver AwsDaemonTaskDefinition#log_driver}.
	// Experimental.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#options AwsDaemonTaskDefinition#options}.
	// Experimental.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// secret_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#secret_option AwsDaemonTaskDefinition#secret_option}
	// Experimental.
	SecretOption interface{} `field:"optional" json:"secretOption" yaml:"secretOption"`
}

