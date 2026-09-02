package awsec2


// Experimental.
type TfSpotInstanceRequest_CapacityReservationSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#capacity_reservation_preference TfSpotInstanceRequest#capacity_reservation_preference}.
	// Experimental.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#capacity_reservation_target TfSpotInstanceRequest#capacity_reservation_target}
	// Experimental.
	CapacityReservationTarget *TfSpotInstanceRequest_CapacityReservationTargetProperty `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

