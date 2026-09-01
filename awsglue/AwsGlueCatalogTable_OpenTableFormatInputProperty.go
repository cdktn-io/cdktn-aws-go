package awsglue


// Experimental.
type AwsGlueCatalogTable_OpenTableFormatInputProperty struct {
	// iceberg_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table#iceberg_input AwsGlueCatalogTable#iceberg_input}
	// Experimental.
	IcebergInput *AwsGlueCatalogTable_IcebergInputProperty `field:"required" json:"icebergInput" yaml:"icebergInput"`
}

