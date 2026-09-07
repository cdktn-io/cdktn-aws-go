package glue


// Experimental.
type AwsCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputSortOrderFieldsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#direction AwsCatalogTable#direction}.
	// Experimental.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#null_order AwsCatalogTable#null_order}.
	// Experimental.
	NullOrder *string `field:"required" json:"nullOrder" yaml:"nullOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#source_id AwsCatalogTable#source_id}.
	// Experimental.
	SourceId *float64 `field:"required" json:"sourceId" yaml:"sourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#transform AwsCatalogTable#transform}.
	// Experimental.
	Transform *string `field:"required" json:"transform" yaml:"transform"`
}

