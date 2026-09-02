package awsglue


// Experimental.
type TfCatalogTable_SchemaIdProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#registry_name TfCatalogTable#registry_name}.
	// Experimental.
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_arn TfCatalogTable#schema_arn}.
	// Experimental.
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_name TfCatalogTable#schema_name}.
	// Experimental.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

