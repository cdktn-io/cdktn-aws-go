package awscloudfront


// Experimental.
type AwsCloudfrontFieldLevelEncryptionProfile_ItemsProperty struct {
	// field_patterns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_profile#field_patterns AwsCloudfrontFieldLevelEncryptionProfile#field_patterns}
	// Experimental.
	FieldPatterns *AwsCloudfrontFieldLevelEncryptionProfile_FieldPatternsProperty `field:"required" json:"fieldPatterns" yaml:"fieldPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_profile#provider_id AwsCloudfrontFieldLevelEncryptionProfile#provider_id}.
	// Experimental.
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_profile#public_key_id AwsCloudfrontFieldLevelEncryptionProfile#public_key_id}.
	// Experimental.
	PublicKeyId *string `field:"required" json:"publicKeyId" yaml:"publicKeyId"`
}

