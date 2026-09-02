package awsec2


// Experimental.
type TfFleet_SpotOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#allocation_strategy TfFleet#allocation_strategy}.
	// Experimental.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#instance_interruption_behavior TfFleet#instance_interruption_behavior}.
	// Experimental.
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#instance_pools_to_use_count TfFleet#instance_pools_to_use_count}.
	// Experimental.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// maintenance_strategies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#maintenance_strategies TfFleet#maintenance_strategies}
	// Experimental.
	MaintenanceStrategies *TfFleet_MaintenanceStrategiesProperty `field:"optional" json:"maintenanceStrategies" yaml:"maintenanceStrategies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#max_total_price TfFleet#max_total_price}.
	// Experimental.
	MaxTotalPrice *string `field:"optional" json:"maxTotalPrice" yaml:"maxTotalPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#min_target_capacity TfFleet#min_target_capacity}.
	// Experimental.
	MinTargetCapacity *float64 `field:"optional" json:"minTargetCapacity" yaml:"minTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#single_availability_zone TfFleet#single_availability_zone}.
	// Experimental.
	SingleAvailabilityZone interface{} `field:"optional" json:"singleAvailabilityZone" yaml:"singleAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#single_instance_type TfFleet#single_instance_type}.
	// Experimental.
	SingleInstanceType interface{} `field:"optional" json:"singleInstanceType" yaml:"singleInstanceType"`
}

