package glue


// Experimental.
type AwsCatalogTable_SkewedInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#skewed_column_names AwsCatalogTable#skewed_column_names}.
	// Experimental.
	SkewedColumnNames *[]*string `field:"optional" json:"skewedColumnNames" yaml:"skewedColumnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#skewed_column_value_location_maps AwsCatalogTable#skewed_column_value_location_maps}.
	// Experimental.
	SkewedColumnValueLocationMaps *map[string]*string `field:"optional" json:"skewedColumnValueLocationMaps" yaml:"skewedColumnValueLocationMaps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#skewed_column_values AwsCatalogTable#skewed_column_values}.
	// Experimental.
	SkewedColumnValues *[]*string `field:"optional" json:"skewedColumnValues" yaml:"skewedColumnValues"`
}

