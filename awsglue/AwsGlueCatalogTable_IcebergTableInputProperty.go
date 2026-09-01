package awsglue


// Experimental.
type AwsGlueCatalogTable_IcebergTableInputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#location AwsGlueCatalogTable#location}.
	// Experimental.
	Location *string `field:"required" json:"location" yaml:"location"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema AwsGlueCatalogTable#schema}
	// Experimental.
	Schema *AwsGlueCatalogTable_SchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// partition_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#partition_spec AwsGlueCatalogTable#partition_spec}
	// Experimental.
	PartitionSpec *AwsGlueCatalogTable_PartitionSpecProperty `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#properties AwsGlueCatalogTable#properties}.
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// sort_order block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#sort_order AwsGlueCatalogTable#sort_order}
	// Experimental.
	SortOrder *AwsGlueCatalogTable_SortOrderProperty `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

