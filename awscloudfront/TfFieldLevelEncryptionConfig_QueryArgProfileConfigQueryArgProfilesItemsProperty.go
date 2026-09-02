package awscloudfront


// Experimental.
type TfFieldLevelEncryptionConfig_QueryArgProfileConfigQueryArgProfilesItemsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#profile_id TfFieldLevelEncryptionConfig#profile_id}.
	// Experimental.
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_field_level_encryption_config#query_arg TfFieldLevelEncryptionConfig#query_arg}.
	// Experimental.
	QueryArg *string `field:"required" json:"queryArg" yaml:"queryArg"`
}

