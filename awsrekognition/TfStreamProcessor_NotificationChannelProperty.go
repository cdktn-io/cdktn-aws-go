package awsrekognition


// Experimental.
type TfStreamProcessor_NotificationChannelProperty struct {
	// The Amazon Resource Number (ARN) of the Amazon Amazon Simple Notification Service topic to which Amazon Rekognition posts the completion status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#sns_topic_arn TfStreamProcessor#sns_topic_arn}
	// Experimental.
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

