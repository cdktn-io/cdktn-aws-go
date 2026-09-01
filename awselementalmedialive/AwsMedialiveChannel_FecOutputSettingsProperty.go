package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_FecOutputSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#column_depth AwsMedialiveChannel#column_depth}.
	// Experimental.
	ColumnDepth *float64 `field:"optional" json:"columnDepth" yaml:"columnDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#include_fec AwsMedialiveChannel#include_fec}.
	// Experimental.
	IncludeFec *string `field:"optional" json:"includeFec" yaml:"includeFec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#row_length AwsMedialiveChannel#row_length}.
	// Experimental.
	RowLength *float64 `field:"optional" json:"rowLength" yaml:"rowLength"`
}

