package awsquicksight


// Experimental.
type AwsQuicksightTheme_SheetProperty struct {
	// tile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#tile AwsQuicksightTheme#tile}
	// Experimental.
	Tile *AwsQuicksightTheme_TileProperty `field:"optional" json:"tile" yaml:"tile"`
	// tile_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#tile_layout AwsQuicksightTheme#tile_layout}
	// Experimental.
	TileLayout *AwsQuicksightTheme_TileLayoutProperty `field:"optional" json:"tileLayout" yaml:"tileLayout"`
}

