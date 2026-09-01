package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_ArchiveOutputSettingsProperty struct {
	// container_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#container_settings AwsMedialiveChannel#container_settings}
	// Experimental.
	ContainerSettings *AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsProperty `field:"optional" json:"containerSettings" yaml:"containerSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#extension AwsMedialiveChannel#extension}.
	// Experimental.
	Extension *string `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name_modifier AwsMedialiveChannel#name_modifier}.
	// Experimental.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

