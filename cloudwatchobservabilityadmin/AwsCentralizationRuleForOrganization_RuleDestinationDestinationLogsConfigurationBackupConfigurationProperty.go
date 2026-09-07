package cloudwatchobservabilityadmin


// Experimental.
type AwsCentralizationRuleForOrganization_RuleDestinationDestinationLogsConfigurationBackupConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#kms_key_arn AwsCentralizationRuleForOrganization#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#region AwsCentralizationRuleForOrganization#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

