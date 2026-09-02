package awslakeformation


// Experimental.
type TfDataCellsFilter_TableDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#database_name TfDataCellsFilter#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#name TfDataCellsFilter#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#table_catalog_id TfDataCellsFilter#table_catalog_id}.
	// Experimental.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#table_name TfDataCellsFilter#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#column_names TfDataCellsFilter#column_names}.
	// Experimental.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// column_wildcard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#column_wildcard TfDataCellsFilter#column_wildcard}
	// Experimental.
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
	// row_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#row_filter TfDataCellsFilter#row_filter}
	// Experimental.
	RowFilter interface{} `field:"optional" json:"rowFilter" yaml:"rowFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_data_cells_filter#version_id TfDataCellsFilter#version_id}.
	// Experimental.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

