package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsProperty struct {
	// m2ts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#m2ts_settings AwsMedialiveChannel#m2ts_settings}
	// Experimental.
	M2TsSettings *AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsProperty `field:"optional" json:"m2TsSettings" yaml:"m2TsSettings"`
	// raw_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#raw_settings AwsMedialiveChannel#raw_settings}
	// Experimental.
	RawSettings *AwsMedialiveChannel_RawSettingsProperty `field:"optional" json:"rawSettings" yaml:"rawSettings"`
}

