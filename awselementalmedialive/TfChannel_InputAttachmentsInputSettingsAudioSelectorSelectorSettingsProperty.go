package awselementalmedialive


// Experimental.
type TfChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty struct {
	// audio_hls_rendition_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_hls_rendition_selection TfChannel#audio_hls_rendition_selection}
	// Experimental.
	AudioHlsRenditionSelection *TfChannel_AudioHlsRenditionSelectionProperty `field:"optional" json:"audioHlsRenditionSelection" yaml:"audioHlsRenditionSelection"`
	// audio_language_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_language_selection TfChannel#audio_language_selection}
	// Experimental.
	AudioLanguageSelection *TfChannel_AudioLanguageSelectionProperty `field:"optional" json:"audioLanguageSelection" yaml:"audioLanguageSelection"`
	// audio_pid_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_pid_selection TfChannel#audio_pid_selection}
	// Experimental.
	AudioPidSelection *TfChannel_AudioPidSelectionProperty `field:"optional" json:"audioPidSelection" yaml:"audioPidSelection"`
	// audio_track_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_track_selection TfChannel#audio_track_selection}
	// Experimental.
	AudioTrackSelection *TfChannel_AudioTrackSelectionProperty `field:"optional" json:"audioTrackSelection" yaml:"audioTrackSelection"`
}

