package awscloudfront


// Experimental.
type AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfileConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#forward_when_query_arg_profile_is_unknown AwsCloudfrontFieldLevelEncryptionConfig#forward_when_query_arg_profile_is_unknown}.
	// Experimental.
	ForwardWhenQueryArgProfileIsUnknown interface{} `field:"required" json:"forwardWhenQueryArgProfileIsUnknown" yaml:"forwardWhenQueryArgProfileIsUnknown"`
	// query_arg_profiles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#query_arg_profiles AwsCloudfrontFieldLevelEncryptionConfig#query_arg_profiles}
	// Experimental.
	QueryArgProfiles *AwsCloudfrontFieldLevelEncryptionConfig_QueryArgProfilesProperty `field:"optional" json:"queryArgProfiles" yaml:"queryArgProfiles"`
}

