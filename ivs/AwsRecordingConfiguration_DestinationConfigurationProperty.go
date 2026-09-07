package ivs


// Experimental.
type AwsRecordingConfiguration_DestinationConfigurationProperty struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#s3 AwsRecordingConfiguration#s3}
	// Experimental.
	S3 *AwsRecordingConfiguration_S3Property `field:"required" json:"s3" yaml:"s3"`
}

