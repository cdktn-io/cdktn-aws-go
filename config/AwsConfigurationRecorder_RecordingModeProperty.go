package config


// Experimental.
type AwsConfigurationRecorder_RecordingModeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#recording_frequency AwsConfigurationRecorder#recording_frequency}.
	// Experimental.
	RecordingFrequency *string `field:"optional" json:"recordingFrequency" yaml:"recordingFrequency"`
	// recording_mode_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#recording_mode_override AwsConfigurationRecorder#recording_mode_override}
	// Experimental.
	RecordingModeOverride *AwsConfigurationRecorder_RecordingModeOverrideProperty `field:"optional" json:"recordingModeOverride" yaml:"recordingModeOverride"`
}

