package awselementalmedialive


// Experimental.
type TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsProperty struct {
	// m2ts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#m2ts_settings TfChannel#m2ts_settings}
	// Experimental.
	M2TsSettings *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsProperty `field:"optional" json:"m2TsSettings" yaml:"m2TsSettings"`
	// raw_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#raw_settings TfChannel#raw_settings}
	// Experimental.
	RawSettings *TfChannel_RawSettingsProperty `field:"optional" json:"rawSettings" yaml:"rawSettings"`
}

