package awsconfig


// Experimental.
type AwsConfigConfigurationRecorder_RecordingModeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#recording_frequency AwsConfigConfigurationRecorder#recording_frequency}.
	// Experimental.
	RecordingFrequency *string `field:"optional" json:"recordingFrequency" yaml:"recordingFrequency"`
	// recording_mode_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#recording_mode_override AwsConfigConfigurationRecorder#recording_mode_override}
	// Experimental.
	RecordingModeOverride *AwsConfigConfigurationRecorder_RecordingModeOverrideProperty `field:"optional" json:"recordingModeOverride" yaml:"recordingModeOverride"`
}

