package awscloudfront


// Experimental.
type AwsCloudfrontFieldLevelEncryptionConfig_ContentTypeProfileConfigContentTypeProfilesItemsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#content_type AwsCloudfrontFieldLevelEncryptionConfig#content_type}.
	// Experimental.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#format AwsCloudfrontFieldLevelEncryptionConfig#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#profile_id AwsCloudfrontFieldLevelEncryptionConfig#profile_id}.
	// Experimental.
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
}

