package glue


// Experimental.
type AwsCatalogTable_SchemaProperty struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#fields AwsCatalogTable#fields}
	// Experimental.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#identifier_field_ids AwsCatalogTable#identifier_field_ids}.
	// Experimental.
	IdentifierFieldIds *[]*float64 `field:"optional" json:"identifierFieldIds" yaml:"identifierFieldIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_id AwsCatalogTable#schema_id}.
	// Experimental.
	SchemaId *float64 `field:"optional" json:"schemaId" yaml:"schemaId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#type AwsCatalogTable#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

