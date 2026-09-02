package awsquicksight


// Experimental.
type TfTheme_SheetProperty struct {
	// tile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#tile TfTheme#tile}
	// Experimental.
	Tile *TfTheme_TileProperty `field:"optional" json:"tile" yaml:"tile"`
	// tile_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#tile_layout TfTheme#tile_layout}
	// Experimental.
	TileLayout *TfTheme_TileLayoutProperty `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

