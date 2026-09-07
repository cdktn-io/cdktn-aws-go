package quicksight


// Experimental.
type AwsDataSet_TagColumnOperationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_name AwsDataSet#column_name}.
	// Experimental.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tags AwsDataSet#tags}
	// Experimental.
	Tags interface{} `field:"required" json:"tags" yaml:"tags"`
}

