package awsglue


// Experimental.
type TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsProperty struct {
	// connection_password_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_data_catalog_encryption_settings#connection_password_encryption TfDataCatalogEncryptionSettings#connection_password_encryption}
	// Experimental.
	ConnectionPasswordEncryption *TfDataCatalogEncryptionSettings_ConnectionPasswordEncryptionProperty `field:"required" json:"connectionPasswordEncryption" yaml:"connectionPasswordEncryption"`
	// encryption_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_data_catalog_encryption_settings#encryption_at_rest TfDataCatalogEncryptionSettings#encryption_at_rest}
	// Experimental.
	EncryptionAtRest *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty `field:"required" json:"encryptionAtRest" yaml:"encryptionAtRest"`
}

