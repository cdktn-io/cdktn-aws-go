package awsconnectcustomerprofiles


// Experimental.
type AwsCustomerprofilesDomain_RuleBasedMatchingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#enabled AwsCustomerprofilesDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// attribute_types_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#attribute_types_selector AwsCustomerprofilesDomain#attribute_types_selector}
	// Experimental.
	AttributeTypesSelector *AwsCustomerprofilesDomain_AttributeTypesSelectorProperty `field:"optional" json:"attributeTypesSelector" yaml:"attributeTypesSelector"`
	// conflict_resolution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#conflict_resolution AwsCustomerprofilesDomain#conflict_resolution}
	// Experimental.
	ConflictResolution *AwsCustomerprofilesDomain_RuleBasedMatchingConflictResolutionProperty `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// exporting_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#exporting_config AwsCustomerprofilesDomain#exporting_config}
	// Experimental.
	ExportingConfig *AwsCustomerprofilesDomain_RuleBasedMatchingExportingConfigProperty `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// matching_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#matching_rules AwsCustomerprofilesDomain#matching_rules}
	// Experimental.
	MatchingRules interface{} `field:"optional" json:"matchingRules" yaml:"matchingRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#max_allowed_rule_level_for_matching AwsCustomerprofilesDomain#max_allowed_rule_level_for_matching}.
	// Experimental.
	MaxAllowedRuleLevelForMatching *float64 `field:"optional" json:"maxAllowedRuleLevelForMatching" yaml:"maxAllowedRuleLevelForMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#max_allowed_rule_level_for_merging AwsCustomerprofilesDomain#max_allowed_rule_level_for_merging}.
	// Experimental.
	MaxAllowedRuleLevelForMerging *float64 `field:"optional" json:"maxAllowedRuleLevelForMerging" yaml:"maxAllowedRuleLevelForMerging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#status AwsCustomerprofilesDomain#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

