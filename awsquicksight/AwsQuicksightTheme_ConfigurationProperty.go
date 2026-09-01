package awsquicksight


// Experimental.
type AwsQuicksightTheme_ConfigurationProperty struct {
	// data_color_palette block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#data_color_palette AwsQuicksightTheme#data_color_palette}
	// Experimental.
	DataColorPalette *AwsQuicksightTheme_DataColorPaletteProperty `field:"optional" json:"dataColorPalette" yaml:"dataColorPalette"`
	// sheet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#sheet AwsQuicksightTheme#sheet}
	// Experimental.
	Sheet *AwsQuicksightTheme_SheetProperty `field:"optional" json:"sheet" yaml:"sheet"`
	// typography block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#typography AwsQuicksightTheme#typography}
	// Experimental.
	Typography *AwsQuicksightTheme_TypographyProperty `field:"optional" json:"typography" yaml:"typography"`
	// ui_color_palette block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#ui_color_palette AwsQuicksightTheme#ui_color_palette}
	// Experimental.
	UiColorPalette *AwsQuicksightTheme_UiColorPaletteProperty `field:"optional" json:"uiColorPalette" yaml:"uiColorPalette"`
}

