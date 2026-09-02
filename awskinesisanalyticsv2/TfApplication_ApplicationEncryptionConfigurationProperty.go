package awskinesisanalyticsv2


// Experimental.
type TfApplication_ApplicationEncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#key_type TfApplication#key_type}.
	// Experimental.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#key_id TfApplication#key_id}.
	// Experimental.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

