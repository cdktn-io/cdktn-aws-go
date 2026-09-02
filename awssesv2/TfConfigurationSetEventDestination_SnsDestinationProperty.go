package awssesv2


// Experimental.
type TfConfigurationSetEventDestination_SnsDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#topic_arn TfConfigurationSetEventDestination#topic_arn}.
	// Experimental.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

