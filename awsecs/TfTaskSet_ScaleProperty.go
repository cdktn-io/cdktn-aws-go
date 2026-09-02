package awsecs


// Experimental.
type TfTaskSet_ScaleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_set#unit TfTaskSet#unit}.
	// Experimental.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_set#value TfTaskSet#value}.
	// Experimental.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

