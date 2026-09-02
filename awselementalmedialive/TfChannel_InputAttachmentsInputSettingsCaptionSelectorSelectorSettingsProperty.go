package awselementalmedialive


// Experimental.
type TfChannel_InputAttachmentsInputSettingsCaptionSelectorSelectorSettingsProperty struct {
	// ancillary_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ancillary_source_settings TfChannel#ancillary_source_settings}
	// Experimental.
	AncillarySourceSettings *TfChannel_AncillarySourceSettingsProperty `field:"optional" json:"ancillarySourceSettings" yaml:"ancillarySourceSettings"`
	// arib_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#arib_source_settings TfChannel#arib_source_settings}
	// Experimental.
	AribSourceSettings *TfChannel_AribSourceSettingsProperty `field:"optional" json:"aribSourceSettings" yaml:"aribSourceSettings"`
	// dvb_sub_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dvb_sub_source_settings TfChannel#dvb_sub_source_settings}
	// Experimental.
	DvbSubSourceSettings *TfChannel_DvbSubSourceSettingsProperty `field:"optional" json:"dvbSubSourceSettings" yaml:"dvbSubSourceSettings"`
	// embedded_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#embedded_source_settings TfChannel#embedded_source_settings}
	// Experimental.
	EmbeddedSourceSettings *TfChannel_EmbeddedSourceSettingsProperty `field:"optional" json:"embeddedSourceSettings" yaml:"embeddedSourceSettings"`
	// scte20_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte20_source_settings TfChannel#scte20_source_settings}
	// Experimental.
	Scte20SourceSettings *TfChannel_Scte20SourceSettingsProperty `field:"optional" json:"scte20SourceSettings" yaml:"scte20SourceSettings"`
	// scte27_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte27_source_settings TfChannel#scte27_source_settings}
	// Experimental.
	Scte27SourceSettings *TfChannel_Scte27SourceSettingsProperty `field:"optional" json:"scte27SourceSettings" yaml:"scte27SourceSettings"`
	// teletext_source_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#teletext_source_settings TfChannel#teletext_source_settings}
	// Experimental.
	TeletextSourceSettings *TfChannel_TeletextSourceSettingsProperty `field:"optional" json:"teletextSourceSettings" yaml:"teletextSourceSettings"`
}

