package awscloudwatchobservabilityadmin


// Experimental.
type TfCentralizationRuleForOrganization_TagPropagationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#destination_role_arn TfCentralizationRuleForOrganization#destination_role_arn}.
	// Experimental.
	DestinationRoleArn *string `field:"required" json:"destinationRoleArn" yaml:"destinationRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_centralization_rule_for_organization#tag_conflict_resolution_strategy TfCentralizationRuleForOrganization#tag_conflict_resolution_strategy}.
	// Experimental.
	TagConflictResolutionStrategy *string `field:"optional" json:"tagConflictResolutionStrategy" yaml:"tagConflictResolutionStrategy"`
}

