package awsivs


// Experimental.
type AwsIvsRecordingConfiguration_DestinationConfigurationProperty struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#s3 AwsIvsRecordingConfiguration#s3}
	// Experimental.
	S3 *AwsIvsRecordingConfiguration_S3Property `field:"required" json:"s3" yaml:"s3"`
}

