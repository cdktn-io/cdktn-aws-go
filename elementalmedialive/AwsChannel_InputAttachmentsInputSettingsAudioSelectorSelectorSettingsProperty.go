package elementalmedialive


// Experimental.
type AwsChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty struct {
	// audio_hls_rendition_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_hls_rendition_selection AwsChannel#audio_hls_rendition_selection}
	// Experimental.
	AudioHlsRenditionSelection *AwsChannel_AudioHlsRenditionSelectionProperty `field:"optional" json:"audioHlsRenditionSelection" yaml:"audioHlsRenditionSelection"`
	// audio_language_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_language_selection AwsChannel#audio_language_selection}
	// Experimental.
	AudioLanguageSelection *AwsChannel_AudioLanguageSelectionProperty `field:"optional" json:"audioLanguageSelection" yaml:"audioLanguageSelection"`
	// audio_pid_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_pid_selection AwsChannel#audio_pid_selection}
	// Experimental.
	AudioPidSelection *AwsChannel_AudioPidSelectionProperty `field:"optional" json:"audioPidSelection" yaml:"audioPidSelection"`
	// audio_track_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_track_selection AwsChannel#audio_track_selection}
	// Experimental.
	AudioTrackSelection *AwsChannel_AudioTrackSelectionProperty `field:"optional" json:"audioTrackSelection" yaml:"audioTrackSelection"`
}

