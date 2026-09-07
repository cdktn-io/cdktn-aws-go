package lakeformation


// Experimental.
type AwsDataCellsFilter_RowFilterProperty struct {
	// all_rows_wildcard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#all_rows_wildcard AwsDataCellsFilter#all_rows_wildcard}
	// Experimental.
	AllRowsWildcard interface{} `field:"optional" json:"allRowsWildcard" yaml:"allRowsWildcard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#filter_expression AwsDataCellsFilter#filter_expression}.
	// Experimental.
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
}

