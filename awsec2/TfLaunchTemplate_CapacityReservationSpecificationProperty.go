package awsec2


// Experimental.
type TfLaunchTemplate_CapacityReservationSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#capacity_reservation_preference TfLaunchTemplate#capacity_reservation_preference}.
	// Experimental.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#capacity_reservation_target TfLaunchTemplate#capacity_reservation_target}
	// Experimental.
	CapacityReservationTarget *TfLaunchTemplate_CapacityReservationTargetProperty `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

