package awsopensearchserverless


// Experimental.
type AwsOpensearchserverlessCollection_EncryptionConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection#aws_owned_key AwsOpensearchserverlessCollection#aws_owned_key}.
	// Experimental.
	AwsOwnedKey interface{} `field:"optional" json:"awsOwnedKey" yaml:"awsOwnedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_collection#kms_key_arn AwsOpensearchserverlessCollection#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

