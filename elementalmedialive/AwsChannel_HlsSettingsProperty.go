package elementalmedialive


// Experimental.
type AwsChannel_HlsSettingsProperty struct {
	// audio_only_hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_only_hls_settings AwsChannel#audio_only_hls_settings}
	// Experimental.
	AudioOnlyHlsSettings *AwsChannel_AudioOnlyHlsSettingsProperty `field:"optional" json:"audioOnlyHlsSettings" yaml:"audioOnlyHlsSettings"`
	// fmp4_hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fmp4_hls_settings AwsChannel#fmp4_hls_settings}
	// Experimental.
	Fmp4HlsSettings *AwsChannel_Fmp4HlsSettingsProperty `field:"optional" json:"fmp4HlsSettings" yaml:"fmp4HlsSettings"`
	// frame_capture_hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_hls_settings AwsChannel#frame_capture_hls_settings}
	// Experimental.
	FrameCaptureHlsSettings *AwsChannel_FrameCaptureHlsSettingsProperty `field:"optional" json:"frameCaptureHlsSettings" yaml:"frameCaptureHlsSettings"`
	// standard_hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#standard_hls_settings AwsChannel#standard_hls_settings}
	// Experimental.
	StandardHlsSettings *AwsChannel_StandardHlsSettingsProperty `field:"optional" json:"standardHlsSettings" yaml:"standardHlsSettings"`
}

