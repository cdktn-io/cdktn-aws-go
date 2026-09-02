package awsglue


// Experimental.
type TfCatalogTable_OpenTableFormatInputIcebergInputIcebergTableInputPartitionSpecFieldsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#name TfCatalogTable#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#source_id TfCatalogTable#source_id}.
	// Experimental.
	SourceId *float64 `field:"required" json:"sourceId" yaml:"sourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#transform TfCatalogTable#transform}.
	// Experimental.
	Transform *string `field:"required" json:"transform" yaml:"transform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#field_id TfCatalogTable#field_id}.
	// Experimental.
	FieldId *float64 `field:"optional" json:"fieldId" yaml:"fieldId"`
}

