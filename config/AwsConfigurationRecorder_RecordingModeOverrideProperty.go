package config


// Experimental.
type AwsConfigurationRecorder_RecordingModeOverrideProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#recording_frequency AwsConfigurationRecorder#recording_frequency}.
	// Experimental.
	RecordingFrequency *string `field:"required" json:"recordingFrequency" yaml:"recordingFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#resource_types AwsConfigurationRecorder#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#description AwsConfigurationRecorder#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

