package rekognition


// Experimental.
type AwsStreamProcessor_OutputProperty struct {
	// kinesis_data_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#kinesis_data_stream AwsStreamProcessor#kinesis_data_stream}
	// Experimental.
	KinesisDataStream interface{} `field:"optional" json:"kinesisDataStream" yaml:"kinesisDataStream"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#s3_destination AwsStreamProcessor#s3_destination}
	// Experimental.
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

