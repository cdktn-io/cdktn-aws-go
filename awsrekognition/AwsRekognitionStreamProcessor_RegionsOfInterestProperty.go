package awsrekognition


// Experimental.
type AwsRekognitionStreamProcessor_RegionsOfInterestProperty struct {
	// bounding_box block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#bounding_box AwsRekognitionStreamProcessor#bounding_box}
	// Experimental.
	BoundingBox interface{} `field:"optional" json:"boundingBox" yaml:"boundingBox"`
	// polygon block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#polygon AwsRekognitionStreamProcessor#polygon}
	// Experimental.
	Polygon interface{} `field:"optional" json:"polygon" yaml:"polygon"`
}

