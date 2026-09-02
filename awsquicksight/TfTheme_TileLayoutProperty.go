package awsquicksight


// Experimental.
type TfTheme_TileLayoutProperty struct {
	// gutter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#gutter TfTheme#gutter}
	// Experimental.
	Gutter *TfTheme_GutterProperty `field:"optional" json:"gutter" yaml:"gutter"`
	// margin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#margin TfTheme#margin}
	// Experimental.
	Margin *TfTheme_MarginProperty `field:"optional" json:"margin" yaml:"margin"`
}

