package awsglue


// Experimental.
type TfCatalogTable_IcebergTableInputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#location TfCatalogTable#location}.
	// Experimental.
	Location *string `field:"required" json:"location" yaml:"location"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#schema TfCatalogTable#schema}
	// Experimental.
	Schema *TfCatalogTable_SchemaProperty `field:"required" json:"schema" yaml:"schema"`
	// partition_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#partition_spec TfCatalogTable#partition_spec}
	// Experimental.
	PartitionSpec *TfCatalogTable_PartitionSpecProperty `field:"optional" json:"partitionSpec" yaml:"partitionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#properties TfCatalogTable#properties}.
	// Experimental.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// sort_order block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#sort_order TfCatalogTable#sort_order}
	// Experimental.
	SortOrder *TfCatalogTable_SortOrderProperty `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

