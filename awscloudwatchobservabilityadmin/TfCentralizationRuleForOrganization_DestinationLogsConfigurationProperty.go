package awscloudwatchobservabilityadmin


// Experimental.
type TfCentralizationRuleForOrganization_DestinationLogsConfigurationProperty struct {
	// backup_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#backup_configuration TfCentralizationRuleForOrganization#backup_configuration}
	// Experimental.
	BackupConfiguration interface{} `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
	// log_group_name_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#log_group_name_configuration TfCentralizationRuleForOrganization#log_group_name_configuration}
	// Experimental.
	LogGroupNameConfiguration interface{} `field:"optional" json:"logGroupNameConfiguration" yaml:"logGroupNameConfiguration"`
	// logs_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#logs_encryption_configuration TfCentralizationRuleForOrganization#logs_encryption_configuration}
	// Experimental.
	LogsEncryptionConfiguration interface{} `field:"optional" json:"logsEncryptionConfiguration" yaml:"logsEncryptionConfiguration"`
	// tag_propagation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#tag_propagation_configuration TfCentralizationRuleForOrganization#tag_propagation_configuration}
	// Experimental.
	TagPropagationConfiguration interface{} `field:"optional" json:"tagPropagationConfiguration" yaml:"tagPropagationConfiguration"`
}

