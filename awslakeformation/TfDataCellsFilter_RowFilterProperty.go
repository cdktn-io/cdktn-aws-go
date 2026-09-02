package awslakeformation


// Experimental.
type TfDataCellsFilter_RowFilterProperty struct {
	// all_rows_wildcard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#all_rows_wildcard TfDataCellsFilter#all_rows_wildcard}
	// Experimental.
	AllRowsWildcard interface{} `field:"optional" json:"allRowsWildcard" yaml:"allRowsWildcard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#filter_expression TfDataCellsFilter#filter_expression}.
	// Experimental.
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
}

