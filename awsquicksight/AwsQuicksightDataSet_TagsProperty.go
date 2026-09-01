package awsquicksight


// Experimental.
type AwsQuicksightDataSet_TagsProperty struct {
	// column_description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_description AwsQuicksightDataSet#column_description}
	// Experimental.
	ColumnDescription *AwsQuicksightDataSet_ColumnDescriptionProperty `field:"optional" json:"columnDescription" yaml:"columnDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_geographic_role AwsQuicksightDataSet#column_geographic_role}.
	// Experimental.
	ColumnGeographicRole *string `field:"optional" json:"columnGeographicRole" yaml:"columnGeographicRole"`
}

