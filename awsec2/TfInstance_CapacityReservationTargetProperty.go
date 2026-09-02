package awsec2


// Experimental.
type TfInstance_CapacityReservationTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#capacity_reservation_id TfInstance#capacity_reservation_id}.
	// Experimental.
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#capacity_reservation_resource_group_arn TfInstance#capacity_reservation_resource_group_arn}.
	// Experimental.
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

