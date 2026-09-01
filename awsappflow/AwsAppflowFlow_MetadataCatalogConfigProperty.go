package awsappflow


// Experimental.
type AwsAppflowFlow_MetadataCatalogConfigProperty struct {
	// glue_data_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#glue_data_catalog AwsAppflowFlow#glue_data_catalog}
	// Experimental.
	GlueDataCatalog *AwsAppflowFlow_GlueDataCatalogProperty `field:"optional" json:"glueDataCatalog" yaml:"glueDataCatalog"`
}

