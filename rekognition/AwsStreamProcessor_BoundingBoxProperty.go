package rekognition


// Experimental.
type AwsStreamProcessor_BoundingBoxProperty struct {
	// Height of the bounding box as a ratio of the overall image height.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#height AwsStreamProcessor#height}
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Left coordinate of the bounding box as a ratio of overall image width.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#left AwsStreamProcessor#left}
	// Experimental.
	Left *float64 `field:"optional" json:"left" yaml:"left"`
	// Top coordinate of the bounding box as a ratio of overall image height.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#top AwsStreamProcessor#top}
	// Experimental.
	Top *float64 `field:"optional" json:"top" yaml:"top"`
	// Width of the bounding box as a ratio of the overall image width.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#width AwsStreamProcessor#width}
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

