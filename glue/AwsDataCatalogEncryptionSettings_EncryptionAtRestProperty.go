package glue


// Experimental.
type AwsDataCatalogEncryptionSettings_EncryptionAtRestProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_data_catalog_encryption_settings#catalog_encryption_mode AwsDataCatalogEncryptionSettings#catalog_encryption_mode}.
	// Experimental.
	CatalogEncryptionMode *string `field:"required" json:"catalogEncryptionMode" yaml:"catalogEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_data_catalog_encryption_settings#catalog_encryption_service_role AwsDataCatalogEncryptionSettings#catalog_encryption_service_role}.
	// Experimental.
	CatalogEncryptionServiceRole *string `field:"optional" json:"catalogEncryptionServiceRole" yaml:"catalogEncryptionServiceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_data_catalog_encryption_settings#sse_aws_kms_key_id AwsDataCatalogEncryptionSettings#sse_aws_kms_key_id}.
	// Experimental.
	SseAwsKmsKeyId *string `field:"optional" json:"sseAwsKmsKeyId" yaml:"sseAwsKmsKeyId"`
}

