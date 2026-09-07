package ivs


// Experimental.
type AwsRecordingConfiguration_ThumbnailConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#recording_mode AwsRecordingConfiguration#recording_mode}.
	// Experimental.
	RecordingMode *string `field:"optional" json:"recordingMode" yaml:"recordingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ivs_recording_configuration#target_interval_seconds AwsRecordingConfiguration#target_interval_seconds}.
	// Experimental.
	TargetIntervalSeconds *float64 `field:"optional" json:"targetIntervalSeconds" yaml:"targetIntervalSeconds"`
}

