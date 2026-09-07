package config


// Experimental.
type AwsConfigurationRecorder_RecordingGroupProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#all_supported AwsConfigurationRecorder#all_supported}.
	// Experimental.
	AllSupported interface{} `field:"optional" json:"allSupported" yaml:"allSupported"`
	// exclusion_by_resource_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#exclusion_by_resource_types AwsConfigurationRecorder#exclusion_by_resource_types}
	// Experimental.
	ExclusionByResourceTypes interface{} `field:"optional" json:"exclusionByResourceTypes" yaml:"exclusionByResourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#include_global_resource_types AwsConfigurationRecorder#include_global_resource_types}.
	// Experimental.
	IncludeGlobalResourceTypes interface{} `field:"optional" json:"includeGlobalResourceTypes" yaml:"includeGlobalResourceTypes"`
	// recording_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#recording_strategy AwsConfigurationRecorder#recording_strategy}
	// Experimental.
	RecordingStrategy interface{} `field:"optional" json:"recordingStrategy" yaml:"recordingStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#resource_types AwsConfigurationRecorder#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

