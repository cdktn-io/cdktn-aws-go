package ecs


// Experimental.
type AwsDaemonTaskDefinition_RestartPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#enabled AwsDaemonTaskDefinition#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#ignored_exit_codes AwsDaemonTaskDefinition#ignored_exit_codes}.
	// Experimental.
	IgnoredExitCodes *[]*float64 `field:"optional" json:"ignoredExitCodes" yaml:"ignoredExitCodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#restart_attempt_period AwsDaemonTaskDefinition#restart_attempt_period}.
	// Experimental.
	RestartAttemptPeriod *float64 `field:"optional" json:"restartAttemptPeriod" yaml:"restartAttemptPeriod"`
}

