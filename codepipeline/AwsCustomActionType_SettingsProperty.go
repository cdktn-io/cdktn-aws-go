package codepipeline


// Experimental.
type AwsCustomActionType_SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#entity_url_template AwsCustomActionType#entity_url_template}.
	// Experimental.
	EntityUrlTemplate *string `field:"optional" json:"entityUrlTemplate" yaml:"entityUrlTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#execution_url_template AwsCustomActionType#execution_url_template}.
	// Experimental.
	ExecutionUrlTemplate *string `field:"optional" json:"executionUrlTemplate" yaml:"executionUrlTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#revision_url_template AwsCustomActionType#revision_url_template}.
	// Experimental.
	RevisionUrlTemplate *string `field:"optional" json:"revisionUrlTemplate" yaml:"revisionUrlTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#third_party_configuration_url AwsCustomActionType#third_party_configuration_url}.
	// Experimental.
	ThirdPartyConfigurationUrl *string `field:"optional" json:"thirdPartyConfigurationUrl" yaml:"thirdPartyConfigurationUrl"`
}

