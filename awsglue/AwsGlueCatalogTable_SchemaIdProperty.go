package awsglue


// Experimental.
type AwsGlueCatalogTable_SchemaIdProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#registry_name AwsGlueCatalogTable#registry_name}.
	// Experimental.
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_arn AwsGlueCatalogTable#schema_arn}.
	// Experimental.
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema_name AwsGlueCatalogTable#schema_name}.
	// Experimental.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

