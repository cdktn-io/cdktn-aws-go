package elementalmedialive


// Experimental.
type AwsChannel_AudioDescriptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_selector_name AwsChannel#audio_selector_name}.
	// Experimental.
	AudioSelectorName *string `field:"required" json:"audioSelectorName" yaml:"audioSelectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name AwsChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// audio_normalization_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_normalization_settings AwsChannel#audio_normalization_settings}
	// Experimental.
	AudioNormalizationSettings *AwsChannel_AudioNormalizationSettingsProperty `field:"optional" json:"audioNormalizationSettings" yaml:"audioNormalizationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_type AwsChannel#audio_type}.
	// Experimental.
	AudioType *string `field:"optional" json:"audioType" yaml:"audioType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_type_control AwsChannel#audio_type_control}.
	// Experimental.
	AudioTypeControl *string `field:"optional" json:"audioTypeControl" yaml:"audioTypeControl"`
	// audio_watermark_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_watermark_settings AwsChannel#audio_watermark_settings}
	// Experimental.
	AudioWatermarkSettings *AwsChannel_AudioWatermarkSettingsProperty `field:"optional" json:"audioWatermarkSettings" yaml:"audioWatermarkSettings"`
	// codec_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#codec_settings AwsChannel#codec_settings}
	// Experimental.
	CodecSettings *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_code AwsChannel#language_code}.
	// Experimental.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#language_code_control AwsChannel#language_code_control}.
	// Experimental.
	LanguageCodeControl *string `field:"optional" json:"languageCodeControl" yaml:"languageCodeControl"`
	// remix_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#remix_settings AwsChannel#remix_settings}
	// Experimental.
	RemixSettings *AwsChannel_RemixSettingsProperty `field:"optional" json:"remixSettings" yaml:"remixSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#stream_name AwsChannel#stream_name}.
	// Experimental.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

