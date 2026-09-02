package awsquicksight


// Experimental.
type TfTheme_ConfigurationProperty struct {
	// data_color_palette block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#data_color_palette TfTheme#data_color_palette}
	// Experimental.
	DataColorPalette *TfTheme_DataColorPaletteProperty `field:"optional" json:"dataColorPalette" yaml:"dataColorPalette"`
	// sheet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#sheet TfTheme#sheet}
	// Experimental.
	Sheet *TfTheme_SheetProperty `field:"optional" json:"sheet" yaml:"sheet"`
	// typography block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#typography TfTheme#typography}
	// Experimental.
	Typography *TfTheme_TypographyProperty `field:"optional" json:"typography" yaml:"typography"`
	// ui_color_palette block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#ui_color_palette TfTheme#ui_color_palette}
	// Experimental.
	UiColorPalette *TfTheme_UiColorPaletteProperty `field:"optional" json:"uiColorPalette" yaml:"uiColorPalette"`
}

