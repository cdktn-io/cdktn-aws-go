package awsquicksight


// Experimental.
type TfDataSet_RenameColumnOperationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_name TfDataSet#column_name}.
	// Experimental.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#new_column_name TfDataSet#new_column_name}.
	// Experimental.
	NewColumnName *string `field:"required" json:"newColumnName" yaml:"newColumnName"`
}

