package glue


// Experimental.
type AwsCatalogTable_RepresentationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#dialect AwsCatalogTable#dialect}.
	// Experimental.
	Dialect *string `field:"optional" json:"dialect" yaml:"dialect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#dialect_version AwsCatalogTable#dialect_version}.
	// Experimental.
	DialectVersion *string `field:"optional" json:"dialectVersion" yaml:"dialectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#validation_connection AwsCatalogTable#validation_connection}.
	// Experimental.
	ValidationConnection *string `field:"optional" json:"validationConnection" yaml:"validationConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_expanded_text AwsCatalogTable#view_expanded_text}.
	// Experimental.
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_original_text AwsCatalogTable#view_original_text}.
	// Experimental.
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

