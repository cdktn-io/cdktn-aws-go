package awsglue


// Experimental.
type TfPartition_SkewedInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#skewed_column_names TfPartition#skewed_column_names}.
	// Experimental.
	SkewedColumnNames *[]*string `field:"optional" json:"skewedColumnNames" yaml:"skewedColumnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#skewed_column_value_location_maps TfPartition#skewed_column_value_location_maps}.
	// Experimental.
	SkewedColumnValueLocationMaps *map[string]*string `field:"optional" json:"skewedColumnValueLocationMaps" yaml:"skewedColumnValueLocationMaps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_partition#skewed_column_values TfPartition#skewed_column_values}.
	// Experimental.
	SkewedColumnValues *[]*string `field:"optional" json:"skewedColumnValues" yaml:"skewedColumnValues"`
}

