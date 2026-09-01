package awsquicksight


// Experimental.
type AwsQuicksightDataSet_TagColumnOperationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_name AwsQuicksightDataSet#column_name}.
	// Experimental.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tags AwsQuicksightDataSet#tags}
	// Experimental.
	Tags interface{} `field:"required" json:"tags" yaml:"tags"`
}

