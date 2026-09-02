package awscloudfront


// Experimental.
type TfFieldLevelEncryptionConfig_ContentTypeProfileConfigProperty struct {
	// content_type_profiles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#content_type_profiles TfFieldLevelEncryptionConfig#content_type_profiles}
	// Experimental.
	ContentTypeProfiles *TfFieldLevelEncryptionConfig_ContentTypeProfilesProperty `field:"required" json:"contentTypeProfiles" yaml:"contentTypeProfiles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#forward_when_content_type_is_unknown TfFieldLevelEncryptionConfig#forward_when_content_type_is_unknown}.
	// Experimental.
	ForwardWhenContentTypeIsUnknown interface{} `field:"required" json:"forwardWhenContentTypeIsUnknown" yaml:"forwardWhenContentTypeIsUnknown"`
}

