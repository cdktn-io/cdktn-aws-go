package awsecs


// Experimental.
type TfTaskSet_CapacityProviderStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_set#capacity_provider TfTaskSet#capacity_provider}.
	// Experimental.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_set#weight TfTaskSet#weight}.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_set#base TfTaskSet#base}.
	// Experimental.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
}

