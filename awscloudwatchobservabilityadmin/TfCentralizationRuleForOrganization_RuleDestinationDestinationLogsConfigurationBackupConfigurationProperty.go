package awscloudwatchobservabilityadmin


// Experimental.
type TfCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#kms_key_arn TfCentralizationRuleForOrganization#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#region TfCentralizationRuleForOrganization#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

