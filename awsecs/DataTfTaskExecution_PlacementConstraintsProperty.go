package awsecs


// Experimental.
type DataTfTaskExecution_PlacementConstraintsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#type DataTfTaskExecution#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecs_task_execution#expression DataTfTaskExecution#expression}.
	// Experimental.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

