package appflow


// Experimental.
type AwsFlow_MetadataCatalogConfigProperty struct {
	// glue_data_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#glue_data_catalog AwsFlow#glue_data_catalog}
	// Experimental.
	GlueDataCatalog *AwsFlow_GlueDataCatalogProperty `field:"optional" json:"glueDataCatalog" yaml:"glueDataCatalog"`
}

