package quicksight


// Experimental.
type AwsDataSet_DataTransformsProperty struct {
	// cast_column_type_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#cast_column_type_operation AwsDataSet#cast_column_type_operation}
	// Experimental.
	CastColumnTypeOperation *AwsDataSet_CastColumnTypeOperationProperty `field:"optional" json:"castColumnTypeOperation" yaml:"castColumnTypeOperation"`
	// create_columns_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#create_columns_operation AwsDataSet#create_columns_operation}
	// Experimental.
	CreateColumnsOperation *AwsDataSet_CreateColumnsOperationProperty `field:"optional" json:"createColumnsOperation" yaml:"createColumnsOperation"`
	// filter_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#filter_operation AwsDataSet#filter_operation}
	// Experimental.
	FilterOperation *AwsDataSet_FilterOperationProperty `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// project_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#project_operation AwsDataSet#project_operation}
	// Experimental.
	ProjectOperation *AwsDataSet_ProjectOperationProperty `field:"optional" json:"projectOperation" yaml:"projectOperation"`
	// rename_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#rename_column_operation AwsDataSet#rename_column_operation}
	// Experimental.
	RenameColumnOperation *AwsDataSet_RenameColumnOperationProperty `field:"optional" json:"renameColumnOperation" yaml:"renameColumnOperation"`
	// tag_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tag_column_operation AwsDataSet#tag_column_operation}
	// Experimental.
	TagColumnOperation *AwsDataSet_TagColumnOperationProperty `field:"optional" json:"tagColumnOperation" yaml:"tagColumnOperation"`
	// untag_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#untag_column_operation AwsDataSet#untag_column_operation}
	// Experimental.
	UntagColumnOperation *AwsDataSet_UntagColumnOperationProperty `field:"optional" json:"untagColumnOperation" yaml:"untagColumnOperation"`
}

