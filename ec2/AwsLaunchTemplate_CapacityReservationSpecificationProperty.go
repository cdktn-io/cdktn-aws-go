package ec2


// Experimental.
type AwsLaunchTemplate_CapacityReservationSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#capacity_reservation_preference AwsLaunchTemplate#capacity_reservation_preference}.
	// Experimental.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// capacity_reservation_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#capacity_reservation_target AwsLaunchTemplate#capacity_reservation_target}
	// Experimental.
	CapacityReservationTarget *AwsLaunchTemplate_CapacityReservationTargetProperty `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

