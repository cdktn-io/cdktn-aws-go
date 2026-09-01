package awsquicksight


// Experimental.
type AwsQuicksightDataSet_LogicalTableMapDataTransformsCreateColumnsOperationColumnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_id AwsQuicksightDataSet#column_id}.
	// Experimental.
	ColumnId *string `field:"required" json:"columnId" yaml:"columnId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_name AwsQuicksightDataSet#column_name}.
	// Experimental.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#expression AwsQuicksightDataSet#expression}.
	// Experimental.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

