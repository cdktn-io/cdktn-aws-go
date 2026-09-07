package rekognition


// Experimental.
type AwsStreamProcessor_DataSharingPreferenceProperty struct {
	// Do you want to share data with Rekognition to improve model performance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rekognition_stream_processor#opt_in AwsStreamProcessor#opt_in}
	// Experimental.
	OptIn interface{} `field:"required" json:"optIn" yaml:"optIn"`
}

