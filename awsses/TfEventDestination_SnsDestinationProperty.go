package awsses


// Experimental.
type TfEventDestination_SnsDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#topic_arn TfEventDestination#topic_arn}.
	// Experimental.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

