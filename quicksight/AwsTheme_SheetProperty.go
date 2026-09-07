package quicksight


// Experimental.
type AwsTheme_SheetProperty struct {
	// tile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#tile AwsTheme#tile}
	// Experimental.
	Tile *AwsTheme_TileProperty `field:"optional" json:"tile" yaml:"tile"`
	// tile_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#tile_layout AwsTheme#tile_layout}
	// Experimental.
	TileLayout *AwsTheme_TileLayoutProperty `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

