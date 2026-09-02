package awsapprunner


// Experimental.
type TfService_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#kms_key TfService#kms_key}.
	// Experimental.
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

