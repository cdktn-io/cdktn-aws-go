package awsquicksight


// Experimental.
type TfDataSet_ColumnGroupsProperty struct {
	// geo_spatial_column_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#geo_spatial_column_group TfDataSet#geo_spatial_column_group}
	// Experimental.
	GeoSpatialColumnGroup *TfDataSet_GeoSpatialColumnGroupProperty `field:"optional" json:"geoSpatialColumnGroup" yaml:"geoSpatialColumnGroup"`
}

