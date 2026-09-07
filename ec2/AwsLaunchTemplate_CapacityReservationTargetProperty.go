package ec2


// Experimental.
type AwsLaunchTemplate_CapacityReservationTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#capacity_reservation_id AwsLaunchTemplate#capacity_reservation_id}.
	// Experimental.
	CapacityReservationId *string `field:"optional" json:"capacityReservationId" yaml:"capacityReservationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#capacity_reservation_resource_group_arn AwsLaunchTemplate#capacity_reservation_resource_group_arn}.
	// Experimental.
	CapacityReservationResourceGroupArn *string `field:"optional" json:"capacityReservationResourceGroupArn" yaml:"capacityReservationResourceGroupArn"`
}

