package awsautoscaling


// Experimental.
type TfGroup_CapacityReservationTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#capacity_reservation_ids TfGroup#capacity_reservation_ids}.
	// Experimental.
	CapacityReservationIds *[]*string `field:"optional" json:"capacityReservationIds" yaml:"capacityReservationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#capacity_reservation_resource_group_arns TfGroup#capacity_reservation_resource_group_arns}.
	// Experimental.
	CapacityReservationResourceGroupArns *[]*string `field:"optional" json:"capacityReservationResourceGroupArns" yaml:"capacityReservationResourceGroupArns"`
}

