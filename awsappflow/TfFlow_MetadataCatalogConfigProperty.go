package awsappflow


// Experimental.
type TfFlow_MetadataCatalogConfigProperty struct {
	// glue_data_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#glue_data_catalog TfFlow#glue_data_catalog}
	// Experimental.
	GlueDataCatalog *TfFlow_GlueDataCatalogProperty `field:"optional" json:"glueDataCatalog" yaml:"glueDataCatalog"`
}

