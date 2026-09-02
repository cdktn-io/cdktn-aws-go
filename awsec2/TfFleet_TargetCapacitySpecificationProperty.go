package awsec2


// Experimental.
type TfFleet_TargetCapacitySpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#default_target_capacity_type TfFleet#default_target_capacity_type}.
	// Experimental.
	DefaultTargetCapacityType *string `field:"required" json:"defaultTargetCapacityType" yaml:"defaultTargetCapacityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#total_target_capacity TfFleet#total_target_capacity}.
	// Experimental.
	TotalTargetCapacity *float64 `field:"required" json:"totalTargetCapacity" yaml:"totalTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#on_demand_target_capacity TfFleet#on_demand_target_capacity}.
	// Experimental.
	OnDemandTargetCapacity *float64 `field:"optional" json:"onDemandTargetCapacity" yaml:"onDemandTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#spot_target_capacity TfFleet#spot_target_capacity}.
	// Experimental.
	SpotTargetCapacity *float64 `field:"optional" json:"spotTargetCapacity" yaml:"spotTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#target_capacity_unit_type TfFleet#target_capacity_unit_type}.
	// Experimental.
	TargetCapacityUnitType *string `field:"optional" json:"targetCapacityUnitType" yaml:"targetCapacityUnitType"`
}

