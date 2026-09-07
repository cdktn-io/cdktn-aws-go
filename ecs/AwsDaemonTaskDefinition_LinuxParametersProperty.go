package ecs


// Experimental.
type AwsDaemonTaskDefinition_LinuxParametersProperty struct {
	// capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#capabilities AwsDaemonTaskDefinition#capabilities}
	// Experimental.
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#device AwsDaemonTaskDefinition#device}
	// Experimental.
	Device interface{} `field:"optional" json:"device" yaml:"device"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#init_process_enabled AwsDaemonTaskDefinition#init_process_enabled}.
	// Experimental.
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// tmpfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#tmpfs AwsDaemonTaskDefinition#tmpfs}
	// Experimental.
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

