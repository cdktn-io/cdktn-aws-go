package awssesv2


// Experimental.
type TfConfigurationSetEventDestination_KinesisFirehoseDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#delivery_stream_arn TfConfigurationSetEventDestination#delivery_stream_arn}.
	// Experimental.
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#iam_role_arn TfConfigurationSetEventDestination#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
}

