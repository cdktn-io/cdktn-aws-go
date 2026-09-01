package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty struct {
	// ancillary_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ancillary_source_settings AwsMedialiveChannel#ancillary_source_settings}
	// Experimental.
	AncillarySourceSettings *AwsMedialiveChannel_AncillarySourceSettingsProperty `field:"optional" json:"ancillarySourceSettings" yaml:"ancillarySourceSettings"`
	// arib_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#arib_source_settings AwsMedialiveChannel#arib_source_settings}
	// Experimental.
	AribSourceSettings *AwsMedialiveChannel_AribSourceSettingsProperty `field:"optional" json:"aribSourceSettings" yaml:"aribSourceSettings"`
	// dvb_sub_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dvb_sub_source_settings AwsMedialiveChannel#dvb_sub_source_settings}
	// Experimental.
	DvbSubSourceSettings *AwsMedialiveChannel_DvbSubSourceSettingsProperty `field:"optional" json:"dvbSubSourceSettings" yaml:"dvbSubSourceSettings"`
	// embedded_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#embedded_source_settings AwsMedialiveChannel#embedded_source_settings}
	// Experimental.
	EmbeddedSourceSettings *AwsMedialiveChannel_EmbeddedSourceSettingsProperty `field:"optional" json:"embeddedSourceSettings" yaml:"embeddedSourceSettings"`
	// scte20_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte20_source_settings AwsMedialiveChannel#scte20_source_settings}
	// Experimental.
	Scte20SourceSettings *AwsMedialiveChannel_Scte20SourceSettingsProperty `field:"optional" json:"scte20SourceSettings" yaml:"scte20SourceSettings"`
	// scte27_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte27_source_settings AwsMedialiveChannel#scte27_source_settings}
	// Experimental.
	Scte27SourceSettings *AwsMedialiveChannel_Scte27SourceSettingsProperty `field:"optional" json:"scte27SourceSettings" yaml:"scte27SourceSettings"`
	// teletext_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#teletext_source_settings AwsMedialiveChannel#teletext_source_settings}
	// Experimental.
	TeletextSourceSettings *AwsMedialiveChannel_TeletextSourceSettingsProperty `field:"optional" json:"teletextSourceSettings" yaml:"teletextSourceSettings"`
}

