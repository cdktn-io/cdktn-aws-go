package awsglue


// Experimental.
type TfCatalogTable_IcebergInputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#metadata_operation TfCatalogTable#metadata_operation}.
	// Experimental.
	MetadataOperation *string `field:"required" json:"metadataOperation" yaml:"metadataOperation"`
	// iceberg_table_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#iceberg_table_input TfCatalogTable#iceberg_table_input}
	// Experimental.
	IcebergTableInput *TfCatalogTable_IcebergTableInputProperty `field:"optional" json:"icebergTableInput" yaml:"icebergTableInput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#version TfCatalogTable#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

