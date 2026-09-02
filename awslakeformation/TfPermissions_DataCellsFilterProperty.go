package awslakeformation


// Experimental.
type TfPermissions_DataCellsFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#database_name TfPermissions#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#name TfPermissions#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#table_catalog_id TfPermissions#table_catalog_id}.
	// Experimental.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_permissions#table_name TfPermissions#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

