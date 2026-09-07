package mq


// Experimental.
type AwsBroker_EncryptionOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#kms_key_id AwsBroker#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker#use_aws_owned_key AwsBroker#use_aws_owned_key}.
	// Experimental.
	UseAwsOwnedKey interface{} `field:"optional" json:"useAwsOwnedKey" yaml:"useAwsOwnedKey"`
}

