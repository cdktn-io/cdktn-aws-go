package awsconfig


// Experimental.
type TfConfigurationRecorder_RecordingGroupProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#all_supported TfConfigurationRecorder#all_supported}.
	// Experimental.
	AllSupported interface{} `field:"optional" json:"allSupported" yaml:"allSupported"`
	// exclusion_by_resource_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#exclusion_by_resource_types TfConfigurationRecorder#exclusion_by_resource_types}
	// Experimental.
	ExclusionByResourceTypes interface{} `field:"optional" json:"exclusionByResourceTypes" yaml:"exclusionByResourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#include_global_resource_types TfConfigurationRecorder#include_global_resource_types}.
	// Experimental.
	IncludeGlobalResourceTypes interface{} `field:"optional" json:"includeGlobalResourceTypes" yaml:"includeGlobalResourceTypes"`
	// recording_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#recording_strategy TfConfigurationRecorder#recording_strategy}
	// Experimental.
	RecordingStrategy interface{} `field:"optional" json:"recordingStrategy" yaml:"recordingStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_configuration_recorder#resource_types TfConfigurationRecorder#resource_types}.
	// Experimental.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

