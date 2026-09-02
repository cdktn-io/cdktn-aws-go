package awsecs


// Experimental.
type TfCapacityProvider_CapacityReservationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#reservation_group_arn TfCapacityProvider#reservation_group_arn}.
	// Experimental.
	ReservationGroupArn *string `field:"optional" json:"reservationGroupArn" yaml:"reservationGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#reservation_preference TfCapacityProvider#reservation_preference}.
	// Experimental.
	ReservationPreference *string `field:"optional" json:"reservationPreference" yaml:"reservationPreference"`
}

