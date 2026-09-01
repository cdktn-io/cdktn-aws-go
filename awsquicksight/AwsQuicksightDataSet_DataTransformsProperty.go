package awsquicksight


// Experimental.
type AwsQuicksightDataSet_DataTransformsProperty struct {
	// cast_column_type_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#cast_column_type_operation AwsQuicksightDataSet#cast_column_type_operation}
	// Experimental.
	CastColumnTypeOperation *AwsQuicksightDataSet_CastColumnTypeOperationProperty `field:"optional" json:"castColumnTypeOperation" yaml:"castColumnTypeOperation"`
	// create_columns_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#create_columns_operation AwsQuicksightDataSet#create_columns_operation}
	// Experimental.
	CreateColumnsOperation *AwsQuicksightDataSet_CreateColumnsOperationProperty `field:"optional" json:"createColumnsOperation" yaml:"createColumnsOperation"`
	// filter_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#filter_operation AwsQuicksightDataSet#filter_operation}
	// Experimental.
	FilterOperation *AwsQuicksightDataSet_FilterOperationProperty `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// project_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#project_operation AwsQuicksightDataSet#project_operation}
	// Experimental.
	ProjectOperation *AwsQuicksightDataSet_ProjectOperationProperty `field:"optional" json:"projectOperation" yaml:"projectOperation"`
	// rename_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#rename_column_operation AwsQuicksightDataSet#rename_column_operation}
	// Experimental.
	RenameColumnOperation *AwsQuicksightDataSet_RenameColumnOperationProperty `field:"optional" json:"renameColumnOperation" yaml:"renameColumnOperation"`
	// tag_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tag_column_operation AwsQuicksightDataSet#tag_column_operation}
	// Experimental.
	TagColumnOperation *AwsQuicksightDataSet_TagColumnOperationProperty `field:"optional" json:"tagColumnOperation" yaml:"tagColumnOperation"`
	// untag_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#untag_column_operation AwsQuicksightDataSet#untag_column_operation}
	// Experimental.
	UntagColumnOperation *AwsQuicksightDataSet_UntagColumnOperationProperty `field:"optional" json:"untagColumnOperation" yaml:"untagColumnOperation"`
}

