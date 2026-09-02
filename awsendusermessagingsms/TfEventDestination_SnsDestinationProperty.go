package awsendusermessagingsms


// Experimental.
type TfEventDestination_SnsDestinationProperty struct {
	// ARN of the Amazon SNS topic that receives the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_event_destination#topic_arn TfEventDestination#topic_arn}
	// Experimental.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

