package ecs


// Experimental.
type AwsCapacityProvider_AutoScalingGroupProviderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#auto_scaling_group_arn AwsCapacityProvider#auto_scaling_group_arn}.
	// Experimental.
	AutoScalingGroupArn *string `field:"required" json:"autoScalingGroupArn" yaml:"autoScalingGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#managed_draining AwsCapacityProvider#managed_draining}.
	// Experimental.
	ManagedDraining *string `field:"optional" json:"managedDraining" yaml:"managedDraining"`
	// managed_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#managed_scaling AwsCapacityProvider#managed_scaling}
	// Experimental.
	ManagedScaling *AwsCapacityProvider_ManagedScalingProperty `field:"optional" json:"managedScaling" yaml:"managedScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#managed_termination_protection AwsCapacityProvider#managed_termination_protection}.
	// Experimental.
	ManagedTerminationProtection *string `field:"optional" json:"managedTerminationProtection" yaml:"managedTerminationProtection"`
}

