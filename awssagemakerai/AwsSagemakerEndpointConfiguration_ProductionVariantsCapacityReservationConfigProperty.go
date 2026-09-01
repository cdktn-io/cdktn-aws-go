package awssagemakerai


// Experimental.
type AwsSagemakerEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#capacity_reservation_preference AwsSagemakerEndpointConfiguration#capacity_reservation_preference}.
	// Experimental.
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#ml_reservation_arn AwsSagemakerEndpointConfiguration#ml_reservation_arn}.
	// Experimental.
	MlReservationArn *string `field:"optional" json:"mlReservationArn" yaml:"mlReservationArn"`
}

