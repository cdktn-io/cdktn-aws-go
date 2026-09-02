package awsautoscaling


// Experimental.
type TfGroup_InstancesDistributionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#on_demand_allocation_strategy TfGroup#on_demand_allocation_strategy}.
	// Experimental.
	OnDemandAllocationStrategy *string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#on_demand_base_capacity TfGroup#on_demand_base_capacity}.
	// Experimental.
	OnDemandBaseCapacity *float64 `field:"optional" json:"onDemandBaseCapacity" yaml:"onDemandBaseCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#on_demand_percentage_above_base_capacity TfGroup#on_demand_percentage_above_base_capacity}.
	// Experimental.
	OnDemandPercentageAboveBaseCapacity *float64 `field:"optional" json:"onDemandPercentageAboveBaseCapacity" yaml:"onDemandPercentageAboveBaseCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#spot_allocation_strategy TfGroup#spot_allocation_strategy}.
	// Experimental.
	SpotAllocationStrategy *string `field:"optional" json:"spotAllocationStrategy" yaml:"spotAllocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#spot_instance_pools TfGroup#spot_instance_pools}.
	// Experimental.
	SpotInstancePools *float64 `field:"optional" json:"spotInstancePools" yaml:"spotInstancePools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#spot_max_price TfGroup#spot_max_price}.
	// Experimental.
	SpotMaxPrice *string `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
}

