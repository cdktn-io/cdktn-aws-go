package ecs


// Experimental.
type AwsCapacityProvider_ManagedScalingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#instance_warmup_period AwsCapacityProvider#instance_warmup_period}.
	// Experimental.
	InstanceWarmupPeriod *float64 `field:"optional" json:"instanceWarmupPeriod" yaml:"instanceWarmupPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#maximum_scaling_step_size AwsCapacityProvider#maximum_scaling_step_size}.
	// Experimental.
	MaximumScalingStepSize *float64 `field:"optional" json:"maximumScalingStepSize" yaml:"maximumScalingStepSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#minimum_scaling_step_size AwsCapacityProvider#minimum_scaling_step_size}.
	// Experimental.
	MinimumScalingStepSize *float64 `field:"optional" json:"minimumScalingStepSize" yaml:"minimumScalingStepSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#status AwsCapacityProvider#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#target_capacity AwsCapacityProvider#target_capacity}.
	// Experimental.
	TargetCapacity *float64 `field:"optional" json:"targetCapacity" yaml:"targetCapacity"`
}

