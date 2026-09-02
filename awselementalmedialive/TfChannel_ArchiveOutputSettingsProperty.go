package awselementalmedialive


// Experimental.
type TfChannel_ArchiveOutputSettingsProperty struct {
	// container_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#container_settings TfChannel#container_settings}
	// Experimental.
	ContainerSettings *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsProperty `field:"optional" json:"containerSettings" yaml:"containerSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#extension TfChannel#extension}.
	// Experimental.
	Extension *string `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name_modifier TfChannel#name_modifier}.
	// Experimental.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

