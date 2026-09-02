package awsec2


// Experimental.
type TfInstance_CapacityReservationSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#capacity_reservation_preference TfInstance#capacity_reservation_preference}.
	// Experimental.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#capacity_reservation_target TfInstance#capacity_reservation_target}
	// Experimental.
	CapacityReservationTarget *TfInstance_CapacityReservationTargetProperty `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

