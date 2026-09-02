package awsglue


// Experimental.
type TfCatalogTable_SchemaProperty struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#fields TfCatalogTable#fields}
	// Experimental.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#identifier_field_ids TfCatalogTable#identifier_field_ids}.
	// Experimental.
	IdentifierFieldIds *[]*float64 `field:"optional" json:"identifierFieldIds" yaml:"identifierFieldIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_id TfCatalogTable#schema_id}.
	// Experimental.
	SchemaId *float64 `field:"optional" json:"schemaId" yaml:"schemaId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#type TfCatalogTable#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

