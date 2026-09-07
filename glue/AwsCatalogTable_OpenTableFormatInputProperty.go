package glue


// Experimental.
type AwsCatalogTable_OpenTableFormatInputProperty struct {
	// iceberg_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#iceberg_input AwsCatalogTable#iceberg_input}
	// Experimental.
	IcebergInput *AwsCatalogTable_IcebergInputProperty `field:"required" json:"icebergInput" yaml:"icebergInput"`
}

