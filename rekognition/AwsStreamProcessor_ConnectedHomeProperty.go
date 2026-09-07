package rekognition


// Experimental.
type AwsStreamProcessor_ConnectedHomeProperty struct {
	// Specifies what you want to detect in the video, such as people, packages, or pets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#labels AwsStreamProcessor#labels}
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The minimum confidence required to label an object in the video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#min_confidence AwsStreamProcessor#min_confidence}
	// Experimental.
	MinConfidence *float64 `field:"optional" json:"minConfidence" yaml:"minConfidence"`
}

