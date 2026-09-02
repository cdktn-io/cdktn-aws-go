package awsbedrockagentcore


// Experimental.
type TfTokenVaultCmk_KmsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_token_vault_cmk#key_type TfTokenVaultCmk#key_type}.
	// Experimental.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_token_vault_cmk#kms_key_arn TfTokenVaultCmk#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

