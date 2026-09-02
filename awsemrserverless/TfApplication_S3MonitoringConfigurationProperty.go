package awsemrserverless


// Experimental.
type TfApplication_S3MonitoringConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#encryption_key_arn TfApplication#encryption_key_arn}.
	// Experimental.
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#log_uri TfApplication#log_uri}.
	// Experimental.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
}

