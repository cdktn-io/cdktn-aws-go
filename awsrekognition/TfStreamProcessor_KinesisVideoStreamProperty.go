package awsrekognition


// Experimental.
type TfStreamProcessor_KinesisVideoStreamProperty struct {
	// ARN of the Kinesis video stream stream that streams the source video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#arn TfStreamProcessor#arn}
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

