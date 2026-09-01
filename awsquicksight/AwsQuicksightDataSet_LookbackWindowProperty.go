package awsquicksight


// Experimental.
type AwsQuicksightDataSet_LookbackWindowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#column_name AwsQuicksightDataSet#column_name}.
	// Experimental.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#size AwsQuicksightDataSet#size}.
	// Experimental.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#size_unit AwsQuicksightDataSet#size_unit}.
	// Experimental.
	SizeUnit *string `field:"required" json:"sizeUnit" yaml:"sizeUnit"`
}

