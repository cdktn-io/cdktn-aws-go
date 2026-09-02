package awsglue


// Experimental.
type TfCatalog_DataLakeAccessPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#catalog_type TfCatalog#catalog_type}.
	// Experimental.
	CatalogType *string `field:"optional" json:"catalogType" yaml:"catalogType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#data_lake_access TfCatalog#data_lake_access}.
	// Experimental.
	DataLakeAccess interface{} `field:"optional" json:"dataLakeAccess" yaml:"dataLakeAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#data_transfer_role TfCatalog#data_transfer_role}.
	// Experimental.
	DataTransferRole *string `field:"optional" json:"dataTransferRole" yaml:"dataTransferRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_catalog#kms_key TfCatalog#kms_key}.
	// Experimental.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

