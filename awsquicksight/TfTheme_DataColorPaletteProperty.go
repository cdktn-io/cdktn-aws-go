package awsquicksight


// Experimental.
type TfTheme_DataColorPaletteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#colors TfTheme#colors}.
	// Experimental.
	Colors *[]*string `field:"optional" json:"colors" yaml:"colors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#empty_fill_color TfTheme#empty_fill_color}.
	// Experimental.
	EmptyFillColor *string `field:"optional" json:"emptyFillColor" yaml:"emptyFillColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#min_max_gradient TfTheme#min_max_gradient}.
	// Experimental.
	MinMaxGradient *[]*string `field:"optional" json:"minMaxGradient" yaml:"minMaxGradient"`
}

