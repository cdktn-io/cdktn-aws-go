package rekognition


// Experimental.
type AwsStreamProcessor_KinesisVideoStreamProperty struct {
	// ARN of the Kinesis video stream stream that streams the source video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#arn AwsStreamProcessor#arn}
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

