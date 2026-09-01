package awsconfig


// Experimental.
type AwsConfigConfigurationRecorder_RecordingModeOverrideProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#recording_frequency AwsConfigConfigurationRecorder#recording_frequency}.
	// Experimental.
	RecordingFrequency *string `field:"required" json:"recordingFrequency" yaml:"recordingFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#resource_types AwsConfigConfigurationRecorder#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#description AwsConfigConfigurationRecorder#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

