package awsecs


// Experimental.
type DataTfTaskExecution_ContainerOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#name DataTfTaskExecution#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#command DataTfTaskExecution#command}.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#cpu DataTfTaskExecution#cpu}.
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#environment DataTfTaskExecution#environment}
	// Experimental.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#memory DataTfTaskExecution#memory}.
	// Experimental.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#memory_reservation DataTfTaskExecution#memory_reservation}.
	// Experimental.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// resource_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#resource_requirements DataTfTaskExecution#resource_requirements}
	// Experimental.
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
}

