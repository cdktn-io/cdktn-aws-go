package awsquicksight


// Experimental.
type TfDataSet_CastColumnTypeOperationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_name TfDataSet#column_name}.
	// Experimental.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#new_column_type TfDataSet#new_column_type}.
	// Experimental.
	NewColumnType *string `field:"required" json:"newColumnType" yaml:"newColumnType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#format TfDataSet#format}.
	// Experimental.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

