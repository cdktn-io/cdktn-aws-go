package awsathena


// Experimental.
type TfWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#encryption_option TfWorkgroup#encryption_option}.
	// Experimental.
	EncryptionOption *string `field:"optional" json:"encryptionOption" yaml:"encryptionOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#kms_key_arn TfWorkgroup#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

