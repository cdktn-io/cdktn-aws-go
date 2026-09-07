package cloudwatchobservabilityadmin


// Experimental.
type AwsCentralizationRuleForOrganization_SourceLogsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#encrypted_log_group_strategy AwsCentralizationRuleForOrganization#encrypted_log_group_strategy}.
	// Experimental.
	EncryptedLogGroupStrategy *string `field:"required" json:"encryptedLogGroupStrategy" yaml:"encryptedLogGroupStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#data_source_selection_criteria AwsCentralizationRuleForOrganization#data_source_selection_criteria}.
	// Experimental.
	DataSourceSelectionCriteria *string `field:"optional" json:"dataSourceSelectionCriteria" yaml:"dataSourceSelectionCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#log_group_selection_criteria AwsCentralizationRuleForOrganization#log_group_selection_criteria}.
	// Experimental.
	LogGroupSelectionCriteria *string `field:"optional" json:"logGroupSelectionCriteria" yaml:"logGroupSelectionCriteria"`
}

