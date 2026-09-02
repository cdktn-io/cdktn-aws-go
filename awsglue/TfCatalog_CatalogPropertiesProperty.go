package awsglue


// Experimental.
type TfCatalog_CatalogPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#custom_properties TfCatalog#custom_properties}.
	// Experimental.
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
	// data_lake_access_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#data_lake_access_properties TfCatalog#data_lake_access_properties}
	// Experimental.
	DataLakeAccessProperties interface{} `field:"optional" json:"dataLakeAccessProperties" yaml:"dataLakeAccessProperties"`
	// iceberg_optimization_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#iceberg_optimization_properties TfCatalog#iceberg_optimization_properties}
	// Experimental.
	IcebergOptimizationProperties interface{} `field:"optional" json:"icebergOptimizationProperties" yaml:"icebergOptimizationProperties"`
}

