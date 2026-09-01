package awskendra


// Experimental.
type AwsKendraExperience_ConfigurationProperty struct {
	// content_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_experience#content_source_configuration AwsKendraExperience#content_source_configuration}
	// Experimental.
	ContentSourceConfiguration *AwsKendraExperience_ContentSourceConfigurationProperty `field:"optional" json:"contentSourceConfiguration" yaml:"contentSourceConfiguration"`
	// user_identity_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_experience#user_identity_configuration AwsKendraExperience#user_identity_configuration}
	// Experimental.
	UserIdentityConfiguration *AwsKendraExperience_UserIdentityConfigurationProperty `field:"optional" json:"userIdentityConfiguration" yaml:"userIdentityConfiguration"`
}

