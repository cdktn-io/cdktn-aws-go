package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_DestinationSettingsProperty struct {
	// arib_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#arib_destination_settings AwsMedialiveChannel#arib_destination_settings}
	// Experimental.
	AribDestinationSettings *AwsMedialiveChannel_AribDestinationSettingsProperty `field:"optional" json:"aribDestinationSettings" yaml:"aribDestinationSettings"`
	// burn_in_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#burn_in_destination_settings AwsMedialiveChannel#burn_in_destination_settings}
	// Experimental.
	BurnInDestinationSettings *AwsMedialiveChannel_BurnInDestinationSettingsProperty `field:"optional" json:"burnInDestinationSettings" yaml:"burnInDestinationSettings"`
	// dvb_sub_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dvb_sub_destination_settings AwsMedialiveChannel#dvb_sub_destination_settings}
	// Experimental.
	DvbSubDestinationSettings *AwsMedialiveChannel_DvbSubDestinationSettingsProperty `field:"optional" json:"dvbSubDestinationSettings" yaml:"dvbSubDestinationSettings"`
	// ebu_tt_d_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ebu_tt_d_destination_settings AwsMedialiveChannel#ebu_tt_d_destination_settings}
	// Experimental.
	EbuTtDDestinationSettings *AwsMedialiveChannel_EbuTtDDestinationSettingsProperty `field:"optional" json:"ebuTtDDestinationSettings" yaml:"ebuTtDDestinationSettings"`
	// embedded_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#embedded_destination_settings AwsMedialiveChannel#embedded_destination_settings}
	// Experimental.
	EmbeddedDestinationSettings *AwsMedialiveChannel_EmbeddedDestinationSettingsProperty `field:"optional" json:"embeddedDestinationSettings" yaml:"embeddedDestinationSettings"`
	// embedded_plus_scte20_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#embedded_plus_scte20_destination_settings AwsMedialiveChannel#embedded_plus_scte20_destination_settings}
	// Experimental.
	EmbeddedPlusScte20DestinationSettings *AwsMedialiveChannel_EmbeddedPlusScte20DestinationSettingsProperty `field:"optional" json:"embeddedPlusScte20DestinationSettings" yaml:"embeddedPlusScte20DestinationSettings"`
	// rtmp_caption_info_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rtmp_caption_info_destination_settings AwsMedialiveChannel#rtmp_caption_info_destination_settings}
	// Experimental.
	RtmpCaptionInfoDestinationSettings *AwsMedialiveChannel_RtmpCaptionInfoDestinationSettingsProperty `field:"optional" json:"rtmpCaptionInfoDestinationSettings" yaml:"rtmpCaptionInfoDestinationSettings"`
	// scte20_plus_embedded_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte20_plus_embedded_destination_settings AwsMedialiveChannel#scte20_plus_embedded_destination_settings}
	// Experimental.
	Scte20PlusEmbeddedDestinationSettings *AwsMedialiveChannel_Scte20PlusEmbeddedDestinationSettingsProperty `field:"optional" json:"scte20PlusEmbeddedDestinationSettings" yaml:"scte20PlusEmbeddedDestinationSettings"`
	// scte27_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte27_destination_settings AwsMedialiveChannel#scte27_destination_settings}
	// Experimental.
	Scte27DestinationSettings *AwsMedialiveChannel_Scte27DestinationSettingsProperty `field:"optional" json:"scte27DestinationSettings" yaml:"scte27DestinationSettings"`
	// smpte_tt_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#smpte_tt_destination_settings AwsMedialiveChannel#smpte_tt_destination_settings}
	// Experimental.
	SmpteTtDestinationSettings *AwsMedialiveChannel_SmpteTtDestinationSettingsProperty `field:"optional" json:"smpteTtDestinationSettings" yaml:"smpteTtDestinationSettings"`
	// teletext_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#teletext_destination_settings AwsMedialiveChannel#teletext_destination_settings}
	// Experimental.
	TeletextDestinationSettings *AwsMedialiveChannel_TeletextDestinationSettingsProperty `field:"optional" json:"teletextDestinationSettings" yaml:"teletextDestinationSettings"`
	// ttml_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ttml_destination_settings AwsMedialiveChannel#ttml_destination_settings}
	// Experimental.
	TtmlDestinationSettings *AwsMedialiveChannel_TtmlDestinationSettingsProperty `field:"optional" json:"ttmlDestinationSettings" yaml:"ttmlDestinationSettings"`
	// webvtt_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#webvtt_destination_settings AwsMedialiveChannel#webvtt_destination_settings}
	// Experimental.
	WebvttDestinationSettings *AwsMedialiveChannel_WebvttDestinationSettingsProperty `field:"optional" json:"webvttDestinationSettings" yaml:"webvttDestinationSettings"`
}

