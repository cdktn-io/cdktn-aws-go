package awsquicksight


// Experimental.
type TfDataSet_LogicalTableMapDataTransformsCreateColumnsOperationColumnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_id TfDataSet#column_id}.
	// Experimental.
	ColumnId *string `field:"required" json:"columnId" yaml:"columnId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_name TfDataSet#column_name}.
	// Experimental.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#expression TfDataSet#expression}.
	// Experimental.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

