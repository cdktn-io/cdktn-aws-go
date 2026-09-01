package awsquicksight


// Experimental.
type AwsQuicksightDataSet_PhysicalTableMapProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#physical_table_map_id AwsQuicksightDataSet#physical_table_map_id}.
	// Experimental.
	PhysicalTableMapId *string `field:"required" json:"physicalTableMapId" yaml:"physicalTableMapId"`
	// custom_sql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#custom_sql AwsQuicksightDataSet#custom_sql}
	// Experimental.
	CustomSql *AwsQuicksightDataSet_CustomSqlProperty `field:"optional" json:"customSql" yaml:"customSql"`
	// relational_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#relational_table AwsQuicksightDataSet#relational_table}
	// Experimental.
	RelationalTable *AwsQuicksightDataSet_RelationalTableProperty `field:"optional" json:"relationalTable" yaml:"relationalTable"`
	// s3_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#s3_source AwsQuicksightDataSet#s3_source}
	// Experimental.
	S3Source *AwsQuicksightDataSet_S3SourceProperty `field:"optional" json:"s3Source" yaml:"s3Source"`
}

