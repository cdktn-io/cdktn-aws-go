package quicksight


// Experimental.
type AwsDataSet_TagsProperty struct {
	// column_description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_description AwsDataSet#column_description}
	// Experimental.
	ColumnDescription *AwsDataSet_ColumnDescriptionProperty `field:"optional" json:"columnDescription" yaml:"columnDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_geographic_role AwsDataSet#column_geographic_role}.
	// Experimental.
	ColumnGeographicRole *string `field:"optional" json:"columnGeographicRole" yaml:"columnGeographicRole"`
}

