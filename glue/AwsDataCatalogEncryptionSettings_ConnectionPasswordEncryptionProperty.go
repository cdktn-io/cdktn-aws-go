package glue


// Experimental.
type AwsDataCatalogEncryptionSettings_ConnectionPasswordEncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_data_catalog_encryption_settings#return_connection_password_encrypted AwsDataCatalogEncryptionSettings#return_connection_password_encrypted}.
	// Experimental.
	ReturnConnectionPasswordEncrypted interface{} `field:"required" json:"returnConnectionPasswordEncrypted" yaml:"returnConnectionPasswordEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_data_catalog_encryption_settings#aws_kms_key_id AwsDataCatalogEncryptionSettings#aws_kms_key_id}.
	// Experimental.
	AwsKmsKeyId *string `field:"optional" json:"awsKmsKeyId" yaml:"awsKmsKeyId"`
}

