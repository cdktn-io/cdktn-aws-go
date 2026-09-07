package glue


// Experimental.
type AwsCatalogTable_SortColumnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#column AwsCatalogTable#column}.
	// Experimental.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#sort_order AwsCatalogTable#sort_order}.
	// Experimental.
	SortOrder *float64 `field:"required" json:"sortOrder" yaml:"sortOrder"`
}

