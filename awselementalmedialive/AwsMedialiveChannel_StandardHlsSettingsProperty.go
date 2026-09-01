package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_StandardHlsSettingsProperty struct {
	// m3u8_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#m3u8_settings AwsMedialiveChannel#m3u8_settings}
	// Experimental.
	M3U8Settings *AwsMedialiveChannel_M3u8SettingsProperty `field:"required" json:"m3U8Settings" yaml:"m3U8Settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_rendition_sets AwsMedialiveChannel#audio_rendition_sets}.
	// Experimental.
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
}

