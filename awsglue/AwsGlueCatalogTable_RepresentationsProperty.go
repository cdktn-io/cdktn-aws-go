package awsglue


// Experimental.
type AwsGlueCatalogTable_RepresentationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#dialect AwsGlueCatalogTable#dialect}.
	// Experimental.
	Dialect *string `field:"optional" json:"dialect" yaml:"dialect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#dialect_version AwsGlueCatalogTable#dialect_version}.
	// Experimental.
	DialectVersion *string `field:"optional" json:"dialectVersion" yaml:"dialectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#validation_connection AwsGlueCatalogTable#validation_connection}.
	// Experimental.
	ValidationConnection *string `field:"optional" json:"validationConnection" yaml:"validationConnection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_expanded_text AwsGlueCatalogTable#view_expanded_text}.
	// Experimental.
	ViewExpandedText *string `field:"optional" json:"viewExpandedText" yaml:"viewExpandedText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#view_original_text AwsGlueCatalogTable#view_original_text}.
	// Experimental.
	ViewOriginalText *string `field:"optional" json:"viewOriginalText" yaml:"viewOriginalText"`
}

