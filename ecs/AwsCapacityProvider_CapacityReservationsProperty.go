package ecs


// Experimental.
type AwsCapacityProvider_CapacityReservationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#reservation_group_arn AwsCapacityProvider#reservation_group_arn}.
	// Experimental.
	ReservationGroupArn *string `field:"optional" json:"reservationGroupArn" yaml:"reservationGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#reservation_preference AwsCapacityProvider#reservation_preference}.
	// Experimental.
	ReservationPreference *string `field:"optional" json:"reservationPreference" yaml:"reservationPreference"`
}

