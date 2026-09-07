package cloudfront


// Experimental.
type AwsFieldLevelEncryptionConfig_QueryArgProfileConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#forward_when_query_arg_profile_is_unknown AwsFieldLevelEncryptionConfig#forward_when_query_arg_profile_is_unknown}.
	// Experimental.
	ForwardWhenQueryArgProfileIsUnknown interface{} `field:"required" json:"forwardWhenQueryArgProfileIsUnknown" yaml:"forwardWhenQueryArgProfileIsUnknown"`
	// query_arg_profiles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#query_arg_profiles AwsFieldLevelEncryptionConfig#query_arg_profiles}
	// Experimental.
	QueryArgProfiles *AwsFieldLevelEncryptionConfig_QueryArgProfilesProperty `field:"optional" json:"queryArgProfiles" yaml:"queryArgProfiles"`
}

