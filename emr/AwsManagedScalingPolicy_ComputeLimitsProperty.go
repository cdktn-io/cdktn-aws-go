package emr


// Experimental.
type AwsManagedScalingPolicy_ComputeLimitsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_managed_scaling_policy#maximum_capacity_units AwsManagedScalingPolicy#maximum_capacity_units}.
	// Experimental.
	MaximumCapacityUnits *float64 `field:"required" json:"maximumCapacityUnits" yaml:"maximumCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_managed_scaling_policy#minimum_capacity_units AwsManagedScalingPolicy#minimum_capacity_units}.
	// Experimental.
	MinimumCapacityUnits *float64 `field:"required" json:"minimumCapacityUnits" yaml:"minimumCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_managed_scaling_policy#unit_type AwsManagedScalingPolicy#unit_type}.
	// Experimental.
	UnitType *string `field:"required" json:"unitType" yaml:"unitType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_managed_scaling_policy#maximum_core_capacity_units AwsManagedScalingPolicy#maximum_core_capacity_units}.
	// Experimental.
	MaximumCoreCapacityUnits *float64 `field:"optional" json:"maximumCoreCapacityUnits" yaml:"maximumCoreCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_managed_scaling_policy#maximum_ondemand_capacity_units AwsManagedScalingPolicy#maximum_ondemand_capacity_units}.
	// Experimental.
	MaximumOndemandCapacityUnits *float64 `field:"optional" json:"maximumOndemandCapacityUnits" yaml:"maximumOndemandCapacityUnits"`
}

