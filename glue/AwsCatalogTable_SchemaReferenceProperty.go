package glue


// Experimental.
type AwsCatalogTable_SchemaReferenceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_version_number AwsCatalogTable#schema_version_number}.
	// Experimental.
	SchemaVersionNumber *float64 `field:"required" json:"schemaVersionNumber" yaml:"schemaVersionNumber"`
	// schema_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_id AwsCatalogTable#schema_id}
	// Experimental.
	SchemaId *AwsCatalogTable_SchemaIdProperty `field:"optional" json:"schemaId" yaml:"schemaId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_version_id AwsCatalogTable#schema_version_id}.
	// Experimental.
	SchemaVersionId *string `field:"optional" json:"schemaVersionId" yaml:"schemaVersionId"`
}

