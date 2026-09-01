package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_AudioDescriptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_selector_name AwsMedialiveChannel#audio_selector_name}.
	// Experimental.
	AudioSelectorName *string `field:"required" json:"audioSelectorName" yaml:"audioSelectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name AwsMedialiveChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// audio_normalization_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_normalization_settings AwsMedialiveChannel#audio_normalization_settings}
	// Experimental.
	AudioNormalizationSettings *AwsMedialiveChannel_AudioNormalizationSettingsProperty `field:"optional" json:"audioNormalizationSettings" yaml:"audioNormalizationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_type AwsMedialiveChannel#audio_type}.
	// Experimental.
	AudioType *string `field:"optional" json:"audioType" yaml:"audioType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_type_control AwsMedialiveChannel#audio_type_control}.
	// Experimental.
	AudioTypeControl *string `field:"optional" json:"audioTypeControl" yaml:"audioTypeControl"`
	// audio_watermark_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_watermark_settings AwsMedialiveChannel#audio_watermark_settings}
	// Experimental.
	AudioWatermarkSettings *AwsMedialiveChannel_AudioWatermarkSettingsProperty `field:"optional" json:"audioWatermarkSettings" yaml:"audioWatermarkSettings"`
	// codec_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#codec_settings AwsMedialiveChannel#codec_settings}
	// Experimental.
	CodecSettings *AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_code AwsMedialiveChannel#language_code}.
	// Experimental.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_code_control AwsMedialiveChannel#language_code_control}.
	// Experimental.
	LanguageCodeControl *string `field:"optional" json:"languageCodeControl" yaml:"languageCodeControl"`
	// remix_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#remix_settings AwsMedialiveChannel#remix_settings}
	// Experimental.
	RemixSettings *AwsMedialiveChannel_RemixSettingsProperty `field:"optional" json:"remixSettings" yaml:"remixSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#stream_name AwsMedialiveChannel#stream_name}.
	// Experimental.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

