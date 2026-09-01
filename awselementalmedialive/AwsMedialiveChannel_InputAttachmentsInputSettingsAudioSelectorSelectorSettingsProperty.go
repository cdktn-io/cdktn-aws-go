package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_InputAttachmentsInputSettingsAudioSelectorSelectorSettingsProperty struct {
	// audio_hls_rendition_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_hls_rendition_selection AwsMedialiveChannel#audio_hls_rendition_selection}
	// Experimental.
	AudioHlsRenditionSelection *AwsMedialiveChannel_AudioHlsRenditionSelectionProperty `field:"optional" json:"audioHlsRenditionSelection" yaml:"audioHlsRenditionSelection"`
	// audio_language_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_language_selection AwsMedialiveChannel#audio_language_selection}
	// Experimental.
	AudioLanguageSelection *AwsMedialiveChannel_AudioLanguageSelectionProperty `field:"optional" json:"audioLanguageSelection" yaml:"audioLanguageSelection"`
	// audio_pid_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_pid_selection AwsMedialiveChannel#audio_pid_selection}
	// Experimental.
	AudioPidSelection *AwsMedialiveChannel_AudioPidSelectionProperty `field:"optional" json:"audioPidSelection" yaml:"audioPidSelection"`
	// audio_track_selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_track_selection AwsMedialiveChannel#audio_track_selection}
	// Experimental.
	AudioTrackSelection *AwsMedialiveChannel_AudioTrackSelectionProperty `field:"optional" json:"audioTrackSelection" yaml:"audioTrackSelection"`
}

