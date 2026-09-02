package awskendra


// Experimental.
type TfExperience_ConfigurationProperty struct {
	// content_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_experience#content_source_configuration TfExperience#content_source_configuration}
	// Experimental.
	ContentSourceConfiguration *TfExperience_ContentSourceConfigurationProperty `field:"optional" json:"contentSourceConfiguration" yaml:"contentSourceConfiguration"`
	// user_identity_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_experience#user_identity_configuration TfExperience#user_identity_configuration}
	// Experimental.
	UserIdentityConfiguration *TfExperience_UserIdentityConfigurationProperty `field:"optional" json:"userIdentityConfiguration" yaml:"userIdentityConfiguration"`
}

