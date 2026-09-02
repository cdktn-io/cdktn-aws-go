package awsecs


// Experimental.
type TfCapacityProvider_ManagedScalingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#instance_warmup_period TfCapacityProvider#instance_warmup_period}.
	// Experimental.
	InstanceWarmupPeriod *float64 `field:"optional" json:"instanceWarmupPeriod" yaml:"instanceWarmupPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#maximum_scaling_step_size TfCapacityProvider#maximum_scaling_step_size}.
	// Experimental.
	MaximumScalingStepSize *float64 `field:"optional" json:"maximumScalingStepSize" yaml:"maximumScalingStepSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#minimum_scaling_step_size TfCapacityProvider#minimum_scaling_step_size}.
	// Experimental.
	MinimumScalingStepSize *float64 `field:"optional" json:"minimumScalingStepSize" yaml:"minimumScalingStepSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#status TfCapacityProvider#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#target_capacity TfCapacityProvider#target_capacity}.
	// Experimental.
	TargetCapacity *float64 `field:"optional" json:"targetCapacity" yaml:"targetCapacity"`
}

