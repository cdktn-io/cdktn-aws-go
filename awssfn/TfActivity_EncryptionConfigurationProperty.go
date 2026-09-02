package awssfn


// Experimental.
type TfActivity_EncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_activity#kms_data_key_reuse_period_seconds TfActivity#kms_data_key_reuse_period_seconds}.
	// Experimental.
	KmsDataKeyReusePeriodSeconds *float64 `field:"optional" json:"kmsDataKeyReusePeriodSeconds" yaml:"kmsDataKeyReusePeriodSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_activity#kms_key_id TfActivity#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_activity#type TfActivity#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

