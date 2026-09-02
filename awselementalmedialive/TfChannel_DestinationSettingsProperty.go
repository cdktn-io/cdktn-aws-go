package awselementalmedialive


// Experimental.
type TfChannel_DestinationSettingsProperty struct {
	// arib_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#arib_destination_settings TfChannel#arib_destination_settings}
	// Experimental.
	AribDestinationSettings *TfChannel_AribDestinationSettingsProperty `field:"optional" json:"aribDestinationSettings" yaml:"aribDestinationSettings"`
	// burn_in_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#burn_in_destination_settings TfChannel#burn_in_destination_settings}
	// Experimental.
	BurnInDestinationSettings *TfChannel_BurnInDestinationSettingsProperty `field:"optional" json:"burnInDestinationSettings" yaml:"burnInDestinationSettings"`
	// dvb_sub_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dvb_sub_destination_settings TfChannel#dvb_sub_destination_settings}
	// Experimental.
	DvbSubDestinationSettings *TfChannel_DvbSubDestinationSettingsProperty `field:"optional" json:"dvbSubDestinationSettings" yaml:"dvbSubDestinationSettings"`
	// ebu_tt_d_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ebu_tt_d_destination_settings TfChannel#ebu_tt_d_destination_settings}
	// Experimental.
	EbuTtDDestinationSettings *TfChannel_EbuTtDDestinationSettingsProperty `field:"optional" json:"ebuTtDDestinationSettings" yaml:"ebuTtDDestinationSettings"`
	// embedded_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#embedded_destination_settings TfChannel#embedded_destination_settings}
	// Experimental.
	EmbeddedDestinationSettings *TfChannel_EmbeddedDestinationSettingsProperty `field:"optional" json:"embeddedDestinationSettings" yaml:"embeddedDestinationSettings"`
	// embedded_plus_scte20_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#embedded_plus_scte20_destination_settings TfChannel#embedded_plus_scte20_destination_settings}
	// Experimental.
	EmbeddedPlusScte20DestinationSettings *TfChannel_EmbeddedPlusScte20DestinationSettingsProperty `field:"optional" json:"embeddedPlusScte20DestinationSettings" yaml:"embeddedPlusScte20DestinationSettings"`
	// rtmp_caption_info_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rtmp_caption_info_destination_settings TfChannel#rtmp_caption_info_destination_settings}
	// Experimental.
	RtmpCaptionInfoDestinationSettings *TfChannel_RtmpCaptionInfoDestinationSettingsProperty `field:"optional" json:"rtmpCaptionInfoDestinationSettings" yaml:"rtmpCaptionInfoDestinationSettings"`
	// scte20_plus_embedded_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte20_plus_embedded_destination_settings TfChannel#scte20_plus_embedded_destination_settings}
	// Experimental.
	Scte20PlusEmbeddedDestinationSettings *TfChannel_Scte20PlusEmbeddedDestinationSettingsProperty `field:"optional" json:"scte20PlusEmbeddedDestinationSettings" yaml:"scte20PlusEmbeddedDestinationSettings"`
	// scte27_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte27_destination_settings TfChannel#scte27_destination_settings}
	// Experimental.
	Scte27DestinationSettings *TfChannel_Scte27DestinationSettingsProperty `field:"optional" json:"scte27DestinationSettings" yaml:"scte27DestinationSettings"`
	// smpte_tt_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#smpte_tt_destination_settings TfChannel#smpte_tt_destination_settings}
	// Experimental.
	SmpteTtDestinationSettings *TfChannel_SmpteTtDestinationSettingsProperty `field:"optional" json:"smpteTtDestinationSettings" yaml:"smpteTtDestinationSettings"`
	// teletext_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#teletext_destination_settings TfChannel#teletext_destination_settings}
	// Experimental.
	TeletextDestinationSettings *TfChannel_TeletextDestinationSettingsProperty `field:"optional" json:"teletextDestinationSettings" yaml:"teletextDestinationSettings"`
	// ttml_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ttml_destination_settings TfChannel#ttml_destination_settings}
	// Experimental.
	TtmlDestinationSettings *TfChannel_TtmlDestinationSettingsProperty `field:"optional" json:"ttmlDestinationSettings" yaml:"ttmlDestinationSettings"`
	// webvtt_destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#webvtt_destination_settings TfChannel#webvtt_destination_settings}
	// Experimental.
	WebvttDestinationSettings *TfChannel_WebvttDestinationSettingsProperty `field:"optional" json:"webvttDestinationSettings" yaml:"webvttDestinationSettings"`
}

