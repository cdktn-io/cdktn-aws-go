package elementalmedialive


// Experimental.
type AwsChannel_ArchiveOutputSettingsProperty struct {
	// container_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#container_settings AwsChannel#container_settings}
	// Experimental.
	ContainerSettings *AwsChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsProperty `field:"optional" json:"containerSettings" yaml:"containerSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#extension AwsChannel#extension}.
	// Experimental.
	Extension *string `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name_modifier AwsChannel#name_modifier}.
	// Experimental.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

