package awsecs


// Experimental.
type DataTfTaskExecution_EnvironmentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#key DataTfTaskExecution#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#value DataTfTaskExecution#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

