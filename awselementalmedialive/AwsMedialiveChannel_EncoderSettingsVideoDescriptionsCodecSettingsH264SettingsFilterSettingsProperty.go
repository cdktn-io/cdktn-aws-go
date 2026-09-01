package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsProperty struct {
	// temporal_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#temporal_filter_settings AwsMedialiveChannel#temporal_filter_settings}
	// Experimental.
	TemporalFilterSettings *AwsMedialiveChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsTemporalFilterSettingsProperty `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

