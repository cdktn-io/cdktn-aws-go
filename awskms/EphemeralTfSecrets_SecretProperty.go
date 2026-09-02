package awskms


// Experimental.
type EphemeralTfSecrets_SecretProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#name EphemeralTfSecrets#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#payload EphemeralTfSecrets#payload}.
	// Experimental.
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#context EphemeralTfSecrets#context}.
	// Experimental.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#encryption_algorithm EphemeralTfSecrets#encryption_algorithm}.
	// Experimental.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#grant_tokens EphemeralTfSecrets#grant_tokens}.
	// Experimental.
	GrantTokens *[]*string `field:"optional" json:"grantTokens" yaml:"grantTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/kms_secrets#key_id EphemeralTfSecrets#key_id}.
	// Experimental.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

