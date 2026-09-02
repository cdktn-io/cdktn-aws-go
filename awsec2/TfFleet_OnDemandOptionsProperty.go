package awsec2


// Experimental.
type TfFleet_OnDemandOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#allocation_strategy TfFleet#allocation_strategy}.
	// Experimental.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// capacity_reservation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#capacity_reservation_options TfFleet#capacity_reservation_options}
	// Experimental.
	CapacityReservationOptions *TfFleet_CapacityReservationOptionsProperty `field:"optional" json:"capacityReservationOptions" yaml:"capacityReservationOptions"`
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

