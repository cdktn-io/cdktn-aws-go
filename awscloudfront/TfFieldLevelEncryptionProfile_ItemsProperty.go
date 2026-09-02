package awscloudfront


// Experimental.
type TfFieldLevelEncryptionProfile_ItemsProperty struct {
	// field_patterns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_profile#field_patterns TfFieldLevelEncryptionProfile#field_patterns}
	// Experimental.
	FieldPatterns *TfFieldLevelEncryptionProfile_FieldPatternsProperty `field:"required" json:"fieldPatterns" yaml:"fieldPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_profile#provider_id TfFieldLevelEncryptionProfile#provider_id}.
	// Experimental.
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_profile#public_key_id TfFieldLevelEncryptionProfile#public_key_id}.
	// Experimental.
	PublicKeyId *string `field:"required" json:"publicKeyId" yaml:"publicKeyId"`
}

