package quicksight


// Experimental.
type AwsDataSet_PhysicalTableMapProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#physical_table_map_id AwsDataSet#physical_table_map_id}.
	// Experimental.
	PhysicalTableMapId *string `field:"required" json:"physicalTableMapId" yaml:"physicalTableMapId"`
	// custom_sql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#custom_sql AwsDataSet#custom_sql}
	// Experimental.
	CustomSql *AwsDataSet_CustomSqlProperty `field:"optional" json:"customSql" yaml:"customSql"`
	// relational_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#relational_table AwsDataSet#relational_table}
	// Experimental.
	RelationalTable *AwsDataSet_RelationalTableProperty `field:"optional" json:"relationalTable" yaml:"relationalTable"`
	// s3_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#s3_source AwsDataSet#s3_source}
	// Experimental.
	S3Source *AwsDataSet_S3SourceProperty `field:"optional" json:"s3Source" yaml:"s3Source"`
}

