package awslakeformation


// Experimental.
type TfOptIn_DataCellsFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#database_name TfOptIn#database_name}.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#name TfOptIn#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#table_catalog_id TfOptIn#table_catalog_id}.
	// Experimental.
	TableCatalogId *string `field:"optional" json:"tableCatalogId" yaml:"tableCatalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#table_name TfOptIn#table_name}.
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

