package cloudfront


// Experimental.
type AwsFieldLevelEncryptionProfile_ItemsProperty struct {
	// field_patterns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_profile#field_patterns AwsFieldLevelEncryptionProfile#field_patterns}
	// Experimental.
	FieldPatterns *AwsFieldLevelEncryptionProfile_FieldPatternsProperty `field:"required" json:"fieldPatterns" yaml:"fieldPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_profile#provider_id AwsFieldLevelEncryptionProfile#provider_id}.
	// Experimental.
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_profile#public_key_id AwsFieldLevelEncryptionProfile#public_key_id}.
	// Experimental.
	PublicKeyId *string `field:"required" json:"publicKeyId" yaml:"publicKeyId"`
}

