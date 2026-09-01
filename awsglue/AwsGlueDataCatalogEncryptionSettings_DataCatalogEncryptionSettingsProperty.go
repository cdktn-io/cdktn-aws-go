package awsglue


// Experimental.
type AwsGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsProperty struct {
	// connection_password_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_data_catalog_encryption_settings#connection_password_encryption AwsGlueDataCatalogEncryptionSettings#connection_password_encryption}
	// Experimental.
	ConnectionPasswordEncryption *AwsGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryptionProperty `field:"required" json:"connectionPasswordEncryption" yaml:"connectionPasswordEncryption"`
	// encryption_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_data_catalog_encryption_settings#encryption_at_rest AwsGlueDataCatalogEncryptionSettings#encryption_at_rest}
	// Experimental.
	EncryptionAtRest *AwsGlueDataCatalogEncryptionSettings_EncryptionAtRestProperty `field:"required" json:"encryptionAtRest" yaml:"encryptionAtRest"`
}

