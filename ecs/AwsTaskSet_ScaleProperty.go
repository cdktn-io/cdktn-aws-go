package ecs


// Experimental.
type AwsTaskSet_ScaleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_set#unit AwsTaskSet#unit}.
	// Experimental.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_set#value AwsTaskSet#value}.
	// Experimental.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

