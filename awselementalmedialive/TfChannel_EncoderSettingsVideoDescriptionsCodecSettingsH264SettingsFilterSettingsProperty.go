package awselementalmedialive


// Experimental.
type TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsProperty struct {
	// temporal_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#temporal_filter_settings TfChannel#temporal_filter_settings}
	// Experimental.
	TemporalFilterSettings *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsTemporalFilterSettingsProperty `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

