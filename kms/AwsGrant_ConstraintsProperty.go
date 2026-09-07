package kms


// Experimental.
type AwsGrant_ConstraintsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_grant#encryption_context_equals AwsGrant#encryption_context_equals}.
	// Experimental.
	EncryptionContextEquals *map[string]*string `field:"optional" json:"encryptionContextEquals" yaml:"encryptionContextEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_grant#encryption_context_subset AwsGrant#encryption_context_subset}.
	// Experimental.
	EncryptionContextSubset *map[string]*string `field:"optional" json:"encryptionContextSubset" yaml:"encryptionContextSubset"`
}

