package awscloudfront


// Experimental.
type TfFieldLevelEncryptionConfig_ContentTypeProfileConfigContentTypeProfilesItemsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#content_type TfFieldLevelEncryptionConfig#content_type}.
	// Experimental.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#format TfFieldLevelEncryptionConfig#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#profile_id TfFieldLevelEncryptionConfig#profile_id}.
	// Experimental.
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
}

