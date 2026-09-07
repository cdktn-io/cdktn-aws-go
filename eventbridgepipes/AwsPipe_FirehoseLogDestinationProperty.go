package eventbridgepipes


// Experimental.
type AwsPipe_FirehoseLogDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#delivery_stream_arn AwsPipe#delivery_stream_arn}.
	// Experimental.
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
}

