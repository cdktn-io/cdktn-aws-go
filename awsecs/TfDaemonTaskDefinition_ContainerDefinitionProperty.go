package awsecs


// Experimental.
type TfDaemonTaskDefinition_ContainerDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#image TfDaemonTaskDefinition#image}.
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#command TfDaemonTaskDefinition#command}.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#cpu TfDaemonTaskDefinition#cpu}.
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// depends_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#depends_on TfDaemonTaskDefinition#depends_on}
	// Experimental.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#entry_point TfDaemonTaskDefinition#entry_point}.
	// Experimental.
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#environment TfDaemonTaskDefinition#environment}
	// Experimental.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// environment_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#environment_file TfDaemonTaskDefinition#environment_file}
	// Experimental.
	EnvironmentFile interface{} `field:"optional" json:"environmentFile" yaml:"environmentFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#essential TfDaemonTaskDefinition#essential}.
	// Experimental.
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// firelens_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#firelens_configuration TfDaemonTaskDefinition#firelens_configuration}
	// Experimental.
	FirelensConfiguration interface{} `field:"optional" json:"firelensConfiguration" yaml:"firelensConfiguration"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#health_check TfDaemonTaskDefinition#health_check}
	// Experimental.
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#interactive TfDaemonTaskDefinition#interactive}.
	// Experimental.
	Interactive interface{} `field:"optional" json:"interactive" yaml:"interactive"`
	// linux_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#linux_parameters TfDaemonTaskDefinition#linux_parameters}
	// Experimental.
	LinuxParameters interface{} `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#log_configuration TfDaemonTaskDefinition#log_configuration}
	// Experimental.
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#memory TfDaemonTaskDefinition#memory}.
	// Experimental.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#memory_reservation TfDaemonTaskDefinition#memory_reservation}.
	// Experimental.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// mount_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#mount_point TfDaemonTaskDefinition#mount_point}
	// Experimental.
	MountPoint interface{} `field:"optional" json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#name TfDaemonTaskDefinition#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#privileged TfDaemonTaskDefinition#privileged}.
	// Experimental.
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#pseudo_terminal TfDaemonTaskDefinition#pseudo_terminal}.
	// Experimental.
	PseudoTerminal interface{} `field:"optional" json:"pseudoTerminal" yaml:"pseudoTerminal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#readonly_root_filesystem TfDaemonTaskDefinition#readonly_root_filesystem}.
	// Experimental.
	ReadonlyRootFilesystem interface{} `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// repository_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#repository_credentials TfDaemonTaskDefinition#repository_credentials}
	// Experimental.
	RepositoryCredentials interface{} `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// restart_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#restart_policy TfDaemonTaskDefinition#restart_policy}
	// Experimental.
	RestartPolicy interface{} `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#secret TfDaemonTaskDefinition#secret}
	// Experimental.
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#start_timeout TfDaemonTaskDefinition#start_timeout}.
	// Experimental.
	StartTimeout *float64 `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#stop_timeout TfDaemonTaskDefinition#stop_timeout}.
	// Experimental.
	StopTimeout *float64 `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// system_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#system_control TfDaemonTaskDefinition#system_control}
	// Experimental.
	SystemControl interface{} `field:"optional" json:"systemControl" yaml:"systemControl"`
	// ulimit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#ulimit TfDaemonTaskDefinition#ulimit}
	// Experimental.
	Ulimit interface{} `field:"optional" json:"ulimit" yaml:"ulimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#user TfDaemonTaskDefinition#user}.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#working_directory TfDaemonTaskDefinition#working_directory}.
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

