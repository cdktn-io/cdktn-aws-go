package awselementalmedialive


// Experimental.
type TfChannel_AudioSilenceSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_selector_name TfChannel#audio_selector_name}.
	// Experimental.
	AudioSelectorName *string `field:"required" json:"audioSelectorName" yaml:"audioSelectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_silence_threshold_msec TfChannel#audio_silence_threshold_msec}.
	// Experimental.
	AudioSilenceThresholdMsec *float64 `field:"optional" json:"audioSilenceThresholdMsec" yaml:"audioSilenceThresholdMsec"`
}

