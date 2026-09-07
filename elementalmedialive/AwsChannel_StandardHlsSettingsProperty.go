package elementalmedialive


// Experimental.
type AwsChannel_StandardHlsSettingsProperty struct {
	// m3u8_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#m3u8_settings AwsChannel#m3u8_settings}
	// Experimental.
	M3U8Settings *AwsChannel_M3u8SettingsProperty `field:"required" json:"m3U8Settings" yaml:"m3U8Settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_rendition_sets AwsChannel#audio_rendition_sets}.
	// Experimental.
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
}

