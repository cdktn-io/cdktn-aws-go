package rekognition


// Experimental.
type AwsStreamProcessor_SettingsProperty struct {
	// connected_home block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#connected_home AwsStreamProcessor#connected_home}
	// Experimental.
	ConnectedHome interface{} `field:"optional" json:"connectedHome" yaml:"connectedHome"`
	// face_search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#face_search AwsStreamProcessor#face_search}
	// Experimental.
	FaceSearch interface{} `field:"optional" json:"faceSearch" yaml:"faceSearch"`
}

