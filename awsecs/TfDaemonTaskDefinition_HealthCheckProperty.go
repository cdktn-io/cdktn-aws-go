package awsecs


// Experimental.
type TfDaemonTaskDefinition_HealthCheckProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#command TfDaemonTaskDefinition#command}.
	// Experimental.
	Command *[]*string `field:"required" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#interval TfDaemonTaskDefinition#interval}.
	// Experimental.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#retries TfDaemonTaskDefinition#retries}.
	// Experimental.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#start_period TfDaemonTaskDefinition#start_period}.
	// Experimental.
	StartPeriod *float64 `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#timeout TfDaemonTaskDefinition#timeout}.
	// Experimental.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

