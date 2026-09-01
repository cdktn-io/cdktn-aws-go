package awssesv2


// Experimental.
type AwsSesv2ConfigurationSetEventDestination_KinesisFirehoseDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#delivery_stream_arn AwsSesv2ConfigurationSetEventDestination#delivery_stream_arn}.
	// Experimental.
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#iam_role_arn AwsSesv2ConfigurationSetEventDestination#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
}

