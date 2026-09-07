package cloudwatchobservabilityadmin


// Experimental.
type AwsCentralizationRuleForOrganization_LogsEncryptionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#encryption_strategy AwsCentralizationRuleForOrganization#encryption_strategy}.
	// Experimental.
	EncryptionStrategy *string `field:"required" json:"encryptionStrategy" yaml:"encryptionStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#encryption_conflict_resolution_strategy AwsCentralizationRuleForOrganization#encryption_conflict_resolution_strategy}.
	// Experimental.
	EncryptionConflictResolutionStrategy *string `field:"optional" json:"encryptionConflictResolutionStrategy" yaml:"encryptionConflictResolutionStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#encryption_scope AwsCentralizationRuleForOrganization#encryption_scope}.
	// Experimental.
	EncryptionScope *string `field:"optional" json:"encryptionScope" yaml:"encryptionScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#kms_key_arn AwsCentralizationRuleForOrganization#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

