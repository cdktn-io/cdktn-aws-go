package awsglue


// Experimental.
type TfCatalogTable_SortColumnsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#column TfCatalogTable#column}.
	// Experimental.
	Column *string `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#sort_order TfCatalogTable#sort_order}.
	// Experimental.
	SortOrder *float64 `field:"required" json:"sortOrder" yaml:"sortOrder"`
}

