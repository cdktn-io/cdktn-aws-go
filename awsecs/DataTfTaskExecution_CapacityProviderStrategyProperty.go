package awsecs


// Experimental.
type DataTfTaskExecution_CapacityProviderStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#capacity_provider DataTfTaskExecution#capacity_provider}.
	// Experimental.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#base DataTfTaskExecution#base}.
	// Experimental.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#weight DataTfTaskExecution#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

