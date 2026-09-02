package awsathena


// Experimental.
type TfDatabase_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_database#encryption_option TfDatabase#encryption_option}.
	// Experimental.
	EncryptionOption *string `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_database#kms_key TfDatabase#kms_key}.
	// Experimental.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

