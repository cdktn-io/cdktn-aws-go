package awsglue


// Experimental.
type AwsGlueCatalogTableOptimizer_CompactionConfigurationProperty struct {
	// iceberg_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog_table_optimizer#iceberg_configuration AwsGlueCatalogTableOptimizer#iceberg_configuration}
	// Experimental.
	IcebergConfiguration interface{} `field:"optional" json:"icebergConfiguration" yaml:"icebergConfiguration"`
}

