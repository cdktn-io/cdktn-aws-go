package quicksight


// Experimental.
type AwsDataSet_UntagColumnOperationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_name AwsDataSet#column_name}.
	// Experimental.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#tag_names AwsDataSet#tag_names}.
	// Experimental.
	TagNames *[]*string `field:"required" json:"tagNames" yaml:"tagNames"`
}

