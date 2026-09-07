package ecs


// Experimental.
type AwsDaemonTaskDefinition_CapabilitiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#add AwsDaemonTaskDefinition#add}.
	// Experimental.
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#drop AwsDaemonTaskDefinition#drop}.
	// Experimental.
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

