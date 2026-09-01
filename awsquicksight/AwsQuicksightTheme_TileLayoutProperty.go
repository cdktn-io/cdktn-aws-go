package awsquicksight


// Experimental.
type AwsQuicksightTheme_TileLayoutProperty struct {
	// gutter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#gutter AwsQuicksightTheme#gutter}
	// Experimental.
	Gutter *AwsQuicksightTheme_GutterProperty `field:"optional" json:"gutter" yaml:"gutter"`
	// margin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#margin AwsQuicksightTheme#margin}
	// Experimental.
	Margin *AwsQuicksightTheme_MarginProperty `field:"optional" json:"margin" yaml:"margin"`
}

