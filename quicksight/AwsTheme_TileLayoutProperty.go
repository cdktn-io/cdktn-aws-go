package quicksight


// Experimental.
type AwsTheme_TileLayoutProperty struct {
	// gutter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#gutter AwsTheme#gutter}
	// Experimental.
	Gutter *AwsTheme_GutterProperty `field:"optional" json:"gutter" yaml:"gutter"`
	// margin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_theme#margin AwsTheme#margin}
	// Experimental.
	Margin *AwsTheme_MarginProperty `field:"optional" json:"margin" yaml:"margin"`
}

