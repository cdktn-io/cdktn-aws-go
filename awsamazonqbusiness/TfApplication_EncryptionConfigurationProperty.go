package awsamazonqbusiness


// Experimental.
type TfApplication_EncryptionConfigurationProperty struct {
	// The identifier of the AWS KMS key that is used to encrypt your data.
	//
	// Amazon Q doesn't support asymmetric keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#kms_key_id TfApplication#kms_key_id}
	// Experimental.
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
}

