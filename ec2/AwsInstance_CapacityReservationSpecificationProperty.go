package ec2


// Experimental.
type AwsInstance_CapacityReservationSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#capacity_reservation_preference AwsInstance#capacity_reservation_preference}.
	// Experimental.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#capacity_reservation_target AwsInstance#capacity_reservation_target}
	// Experimental.
	CapacityReservationTarget *AwsInstance_CapacityReservationTargetProperty `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

