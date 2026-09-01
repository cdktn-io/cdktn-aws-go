package awssagemakerai


// Experimental.
type AwsSagemakerFeatureGroup_OfflineStoreConfigProperty struct {
	// s3_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#s3_storage_config AwsSagemakerFeatureGroup#s3_storage_config}
	// Experimental.
	S3StorageConfig *AwsSagemakerFeatureGroup_S3StorageConfigProperty `field:"required" json:"s3StorageConfig" yaml:"s3StorageConfig"`
	// data_catalog_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#data_catalog_config AwsSagemakerFeatureGroup#data_catalog_config}
	// Experimental.
	DataCatalogConfig *AwsSagemakerFeatureGroup_DataCatalogConfigProperty `field:"optional" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#disable_glue_table_creation AwsSagemakerFeatureGroup#disable_glue_table_creation}.
	// Experimental.
	DisableGlueTableCreation interface{} `field:"optional" json:"disableGlueTableCreation" yaml:"disableGlueTableCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#table_format AwsSagemakerFeatureGroup#table_format}.
	// Experimental.
	TableFormat *string `field:"optional" json:"tableFormat" yaml:"tableFormat"`
}

