package awselementalmedialive


// Experimental.
type TfChannel_DvbSubDestinationSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#alignment TfChannel#alignment}.
	// Experimental.
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#background_color TfChannel#background_color}.
	// Experimental.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#background_opacity TfChannel#background_opacity}.
	// Experimental.
	BackgroundOpacity *float64 `field:"optional" json:"backgroundOpacity" yaml:"backgroundOpacity"`
	// font block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#font TfChannel#font}
	// Experimental.
	Font *TfChannel_EncoderSettingsCaptionDescriptionsDestinationSettingsDvbSubDestinationSettingsFontProperty `field:"optional" json:"font" yaml:"font"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#font_color TfChannel#font_color}.
	// Experimental.
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#font_opacity TfChannel#font_opacity}.
	// Experimental.
	FontOpacity *float64 `field:"optional" json:"fontOpacity" yaml:"fontOpacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#font_resolution TfChannel#font_resolution}.
	// Experimental.
	FontResolution *float64 `field:"optional" json:"fontResolution" yaml:"fontResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#font_size TfChannel#font_size}.
	// Experimental.
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#outline_color TfChannel#outline_color}.
	// Experimental.
	OutlineColor *string `field:"optional" json:"outlineColor" yaml:"outlineColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#outline_size TfChannel#outline_size}.
	// Experimental.
	OutlineSize *float64 `field:"optional" json:"outlineSize" yaml:"outlineSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#shadow_color TfChannel#shadow_color}.
	// Experimental.
	ShadowColor *string `field:"optional" json:"shadowColor" yaml:"shadowColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#shadow_opacity TfChannel#shadow_opacity}.
	// Experimental.
	ShadowOpacity *float64 `field:"optional" json:"shadowOpacity" yaml:"shadowOpacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#shadow_x_offset TfChannel#shadow_x_offset}.
	// Experimental.
	ShadowXOffset *float64 `field:"optional" json:"shadowXOffset" yaml:"shadowXOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#shadow_y_offset TfChannel#shadow_y_offset}.
	// Experimental.
	ShadowYOffset *float64 `field:"optional" json:"shadowYOffset" yaml:"shadowYOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#teletext_grid_control TfChannel#teletext_grid_control}.
	// Experimental.
	TeletextGridControl *string `field:"optional" json:"teletextGridControl" yaml:"teletextGridControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#x_position TfChannel#x_position}.
	// Experimental.
	XPosition *float64 `field:"optional" json:"xPosition" yaml:"xPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#y_position TfChannel#y_position}.
	// Experimental.
	YPosition *float64 `field:"optional" json:"yPosition" yaml:"yPosition"`
}

