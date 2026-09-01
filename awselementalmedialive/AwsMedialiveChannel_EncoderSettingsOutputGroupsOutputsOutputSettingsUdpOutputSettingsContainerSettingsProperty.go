package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsProperty struct {
	// m2ts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#m2ts_settings AwsMedialiveChannel#m2ts_settings}
	// Experimental.
	M2TsSettings *AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsProperty `field:"optional" json:"m2TsSettings" yaml:"m2TsSettings"`
}

