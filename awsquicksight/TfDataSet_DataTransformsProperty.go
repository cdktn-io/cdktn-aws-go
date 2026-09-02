package awsquicksight


// Experimental.
type TfDataSet_DataTransformsProperty struct {
	// cast_column_type_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#cast_column_type_operation TfDataSet#cast_column_type_operation}
	// Experimental.
	CastColumnTypeOperation *TfDataSet_CastColumnTypeOperationProperty `field:"optional" json:"castColumnTypeOperation" yaml:"castColumnTypeOperation"`
	// create_columns_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#create_columns_operation TfDataSet#create_columns_operation}
	// Experimental.
	CreateColumnsOperation *TfDataSet_CreateColumnsOperationProperty `field:"optional" json:"createColumnsOperation" yaml:"createColumnsOperation"`
	// filter_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#filter_operation TfDataSet#filter_operation}
	// Experimental.
	FilterOperation *TfDataSet_FilterOperationProperty `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// project_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#project_operation TfDataSet#project_operation}
	// Experimental.
	ProjectOperation *TfDataSet_ProjectOperationProperty `field:"optional" json:"projectOperation" yaml:"projectOperation"`
	// rename_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#rename_column_operation TfDataSet#rename_column_operation}
	// Experimental.
	RenameColumnOperation *TfDataSet_RenameColumnOperationProperty `field:"optional" json:"renameColumnOperation" yaml:"renameColumnOperation"`
	// tag_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tag_column_operation TfDataSet#tag_column_operation}
	// Experimental.
	TagColumnOperation *TfDataSet_TagColumnOperationProperty `field:"optional" json:"tagColumnOperation" yaml:"tagColumnOperation"`
	// untag_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#untag_column_operation TfDataSet#untag_column_operation}
	// Experimental.
	UntagColumnOperation *TfDataSet_UntagColumnOperationProperty `field:"optional" json:"untagColumnOperation" yaml:"untagColumnOperation"`
}

