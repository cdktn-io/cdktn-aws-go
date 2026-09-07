package elementalmedialive


// Experimental.
type AwsChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsProperty struct {
	// m2ts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#m2ts_settings AwsChannel#m2ts_settings}
	// Experimental.
	M2TsSettings *AwsChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsProperty `field:"optional" json:"m2TsSettings" yaml:"m2TsSettings"`
	// raw_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#raw_settings AwsChannel#raw_settings}
	// Experimental.
	RawSettings *AwsChannel_RawSettingsProperty `field:"optional" json:"rawSettings" yaml:"rawSettings"`
}

