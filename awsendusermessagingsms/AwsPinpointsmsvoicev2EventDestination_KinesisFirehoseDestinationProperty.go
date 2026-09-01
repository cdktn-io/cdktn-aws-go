package awsendusermessagingsms


// Experimental.
type AwsPinpointsmsvoicev2EventDestination_KinesisFirehoseDestinationProperty struct {
	// ARN of the Amazon Data Firehose delivery stream that receives the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_event_destination#delivery_stream_arn AwsPinpointsmsvoicev2EventDestination#delivery_stream_arn}
	// Experimental.
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// ARN of the IAM role that End User Messaging SMS assumes to write to the delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_event_destination#iam_role_arn AwsPinpointsmsvoicev2EventDestination#iam_role_arn}
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
}

