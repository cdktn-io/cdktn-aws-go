package awselementalmedialive


// Experimental.
type TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsProperty struct {
	// temporal_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#temporal_filter_settings TfChannel#temporal_filter_settings}
	// Experimental.
	TemporalFilterSettings *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsTemporalFilterSettingsProperty `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

