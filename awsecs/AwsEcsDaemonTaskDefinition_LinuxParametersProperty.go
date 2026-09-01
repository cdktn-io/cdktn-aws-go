package awsecs


// Experimental.
type AwsEcsDaemonTaskDefinition_LinuxParametersProperty struct {
	// capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#capabilities AwsEcsDaemonTaskDefinition#capabilities}
	// Experimental.
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#device AwsEcsDaemonTaskDefinition#device}
	// Experimental.
	Device interface{} `field:"optional" json:"device" yaml:"device"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#init_process_enabled AwsEcsDaemonTaskDefinition#init_process_enabled}.
	// Experimental.
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// tmpfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#tmpfs AwsEcsDaemonTaskDefinition#tmpfs}
	// Experimental.
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

