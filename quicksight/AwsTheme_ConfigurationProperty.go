package quicksight


// Experimental.
type AwsTheme_ConfigurationProperty struct {
	// data_color_palette block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#data_color_palette AwsTheme#data_color_palette}
	// Experimental.
	DataColorPalette *AwsTheme_DataColorPaletteProperty `field:"optional" json:"dataColorPalette" yaml:"dataColorPalette"`
	// sheet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#sheet AwsTheme#sheet}
	// Experimental.
	Sheet *AwsTheme_SheetProperty `field:"optional" json:"sheet" yaml:"sheet"`
	// typography block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#typography AwsTheme#typography}
	// Experimental.
	Typography *AwsTheme_TypographyProperty `field:"optional" json:"typography" yaml:"typography"`
	// ui_color_palette block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#ui_color_palette AwsTheme#ui_color_palette}
	// Experimental.
	UiColorPalette *AwsTheme_UiColorPaletteProperty `field:"optional" json:"uiColorPalette" yaml:"uiColorPalette"`
}

