package awslakeformation


// Experimental.
type TfResourceLfTag_TableWithColumnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tag#database_name TfResourceLfTag#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tag#name TfResourceLfTag#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tag#catalog_id TfResourceLfTag#catalog_id}.
	// Experimental.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tag#column_names TfResourceLfTag#column_names}.
	// Experimental.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// column_wildcard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_resource_lf_tag#column_wildcard TfResourceLfTag#column_wildcard}
	// Experimental.
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
}

