package awsrekognition


// Experimental.
type AwsRekognitionStreamProcessor_S3DestinationProperty struct {
	// The name of the Amazon S3 bucket you want to associate with the streaming video project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#bucket AwsRekognitionStreamProcessor#bucket}
	// Experimental.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The prefix value of the location within the bucket that you want the information to be published to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#key_prefix AwsRekognitionStreamProcessor#key_prefix}
	// Experimental.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

