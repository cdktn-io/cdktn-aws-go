package awskms


// Experimental.
type DataTfSecrets_SecretProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secrets#name DataTfSecrets#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secrets#payload DataTfSecrets#payload}.
	// Experimental.
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secrets#context DataTfSecrets#context}.
	// Experimental.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secrets#encryption_algorithm DataTfSecrets#encryption_algorithm}.
	// Experimental.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secrets#grant_tokens DataTfSecrets#grant_tokens}.
	// Experimental.
	GrantTokens *[]*string `field:"optional" json:"grantTokens" yaml:"grantTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/kms_secrets#key_id DataTfSecrets#key_id}.
	// Experimental.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

