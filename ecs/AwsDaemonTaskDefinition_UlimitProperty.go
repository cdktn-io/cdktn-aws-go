package ecs


// Experimental.
type AwsDaemonTaskDefinition_UlimitProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#hard_limit AwsDaemonTaskDefinition#hard_limit}.
	// Experimental.
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#name AwsDaemonTaskDefinition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon_task_definition#soft_limit AwsDaemonTaskDefinition#soft_limit}.
	// Experimental.
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
}

