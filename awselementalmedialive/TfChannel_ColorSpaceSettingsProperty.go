package awselementalmedialive


// Experimental.
type TfChannel_ColorSpaceSettingsProperty struct {
	// color_space_passthrough_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#color_space_passthrough_settings TfChannel#color_space_passthrough_settings}
	// Experimental.
	ColorSpacePassthroughSettings *TfChannel_ColorSpacePassthroughSettingsProperty `field:"optional" json:"colorSpacePassthroughSettings" yaml:"colorSpacePassthroughSettings"`
	// dolby_vision81_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dolby_vision81_settings TfChannel#dolby_vision81_settings}
	// Experimental.
	DolbyVision81Settings *TfChannel_DolbyVision81SettingsProperty `field:"optional" json:"dolbyVision81Settings" yaml:"dolbyVision81Settings"`
	// hdr10_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hdr10_settings TfChannel#hdr10_settings}
	// Experimental.
	Hdr10Settings *TfChannel_Hdr10SettingsProperty `field:"optional" json:"hdr10Settings" yaml:"hdr10Settings"`
	// rec601_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rec601_settings TfChannel#rec601_settings}
	// Experimental.
	Rec601Settings *TfChannel_Rec601SettingsProperty `field:"optional" json:"rec601Settings" yaml:"rec601Settings"`
	// rec709_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rec709_settings TfChannel#rec709_settings}
	// Experimental.
	Rec709Settings *TfChannel_Rec709SettingsProperty `field:"optional" json:"rec709Settings" yaml:"rec709Settings"`
}

