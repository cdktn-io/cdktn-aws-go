package awssagemakerai


// Experimental.
type TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#capacity_reservation_preference TfEndpointConfiguration#capacity_reservation_preference}.
	// Experimental.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#ml_reservation_arn TfEndpointConfiguration#ml_reservation_arn}.
	// Experimental.
	MlReservationArn *string `field:"optional" json:"mlReservationArn" yaml:"mlReservationArn"`
}

