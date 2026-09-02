package awssagemakerai


// Experimental.
type TfFeatureGroup_OfflineStoreConfigProperty struct {
	// s3_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#s3_storage_config TfFeatureGroup#s3_storage_config}
	// Experimental.
	S3StorageConfig *TfFeatureGroup_S3StorageConfigProperty `field:"required" json:"s3StorageConfig" yaml:"s3StorageConfig"`
	// data_catalog_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#data_catalog_config TfFeatureGroup#data_catalog_config}
	// Experimental.
	DataCatalogConfig *TfFeatureGroup_DataCatalogConfigProperty `field:"optional" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#disable_glue_table_creation TfFeatureGroup#disable_glue_table_creation}.
	// Experimental.
	DisableGlueTableCreation interface{} `field:"optional" json:"disableGlueTableCreation" yaml:"disableGlueTableCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#table_format TfFeatureGroup#table_format}.
	// Experimental.
	TableFormat *string `field:"optional" json:"tableFormat" yaml:"tableFormat"`
}

