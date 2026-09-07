package elementalmedialive


// Experimental.
type AwsChannel_FailoverConditionSettingsProperty struct {
	// audio_silence_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_silence_settings AwsChannel#audio_silence_settings}
	// Experimental.
	AudioSilenceSettings *AwsChannel_AudioSilenceSettingsProperty `field:"optional" json:"audioSilenceSettings" yaml:"audioSilenceSettings"`
	// input_loss_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_settings AwsChannel#input_loss_settings}
	// Experimental.
	InputLossSettings *AwsChannel_InputLossSettingsProperty `field:"optional" json:"inputLossSettings" yaml:"inputLossSettings"`
	// video_black_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#video_black_settings AwsChannel#video_black_settings}
	// Experimental.
	VideoBlackSettings *AwsChannel_VideoBlackSettingsProperty `field:"optional" json:"videoBlackSettings" yaml:"videoBlackSettings"`
}

