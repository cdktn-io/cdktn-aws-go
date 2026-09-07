package sagemakerai


// Experimental.
type AwsFeatureGroup_OfflineStoreConfigProperty struct {
	// s3_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#s3_storage_config AwsFeatureGroup#s3_storage_config}
	// Experimental.
	S3StorageConfig *AwsFeatureGroup_S3StorageConfigProperty `field:"required" json:"s3StorageConfig" yaml:"s3StorageConfig"`
	// data_catalog_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#data_catalog_config AwsFeatureGroup#data_catalog_config}
	// Experimental.
	DataCatalogConfig *AwsFeatureGroup_DataCatalogConfigProperty `field:"optional" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#disable_glue_table_creation AwsFeatureGroup#disable_glue_table_creation}.
	// Experimental.
	DisableGlueTableCreation interface{} `field:"optional" json:"disableGlueTableCreation" yaml:"disableGlueTableCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#table_format AwsFeatureGroup#table_format}.
	// Experimental.
	TableFormat *string `field:"optional" json:"tableFormat" yaml:"tableFormat"`
}

