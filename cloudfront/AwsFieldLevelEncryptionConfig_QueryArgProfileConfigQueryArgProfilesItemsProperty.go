package cloudfront


// Experimental.
type AwsFieldLevelEncryptionConfig_QueryArgProfileConfigQueryArgProfilesItemsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#profile_id AwsFieldLevelEncryptionConfig#profile_id}.
	// Experimental.
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#query_arg AwsFieldLevelEncryptionConfig#query_arg}.
	// Experimental.
	QueryArg *string `field:"required" json:"queryArg" yaml:"queryArg"`
}

