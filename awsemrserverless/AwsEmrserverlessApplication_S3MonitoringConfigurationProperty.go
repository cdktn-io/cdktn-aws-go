package awsemrserverless


// Experimental.
type AwsEmrserverlessApplication_S3MonitoringConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#encryption_key_arn AwsEmrserverlessApplication#encryption_key_arn}.
	// Experimental.
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#log_uri AwsEmrserverlessApplication#log_uri}.
	// Experimental.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
}

