package elementalmedialive


// Experimental.
type AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsProperty struct {
	// temporal_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#temporal_filter_settings AwsChannel#temporal_filter_settings}
	// Experimental.
	TemporalFilterSettings *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsTemporalFilterSettingsProperty `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

