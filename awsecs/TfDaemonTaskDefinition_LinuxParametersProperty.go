package awsecs


// Experimental.
type TfDaemonTaskDefinition_LinuxParametersProperty struct {
	// capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#capabilities TfDaemonTaskDefinition#capabilities}
	// Experimental.
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#device TfDaemonTaskDefinition#device}
	// Experimental.
	Device interface{} `field:"optional" json:"device" yaml:"device"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#init_process_enabled TfDaemonTaskDefinition#init_process_enabled}.
	// Experimental.
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// tmpfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#tmpfs TfDaemonTaskDefinition#tmpfs}
	// Experimental.
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

