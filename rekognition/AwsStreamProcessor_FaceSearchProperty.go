package rekognition


// Experimental.
type AwsStreamProcessor_FaceSearchProperty struct {
	// The ID of a collection that contains faces that you want to search for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#collection_id AwsStreamProcessor#collection_id}
	// Experimental.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// Minimum face match confidence score that must be met to return a result for a recognized face.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#face_match_threshold AwsStreamProcessor#face_match_threshold}
	// Experimental.
	FaceMatchThreshold *float64 `field:"optional" json:"faceMatchThreshold" yaml:"faceMatchThreshold"`
}

