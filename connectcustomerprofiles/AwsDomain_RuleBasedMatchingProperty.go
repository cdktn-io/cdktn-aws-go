package connectcustomerprofiles


// Experimental.
type AwsDomain_RuleBasedMatchingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#enabled AwsDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// attribute_types_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#attribute_types_selector AwsDomain#attribute_types_selector}
	// Experimental.
	AttributeTypesSelector *AwsDomain_AttributeTypesSelectorProperty `field:"optional" json:"attributeTypesSelector" yaml:"attributeTypesSelector"`
	// conflict_resolution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#conflict_resolution AwsDomain#conflict_resolution}
	// Experimental.
	ConflictResolution *AwsDomain_RuleBasedMatchingConflictResolutionProperty `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// exporting_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#exporting_config AwsDomain#exporting_config}
	// Experimental.
	ExportingConfig *AwsDomain_RuleBasedMatchingExportingConfigProperty `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// matching_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#matching_rules AwsDomain#matching_rules}
	// Experimental.
	MatchingRules interface{} `field:"optional" json:"matchingRules" yaml:"matchingRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#max_allowed_rule_level_for_matching AwsDomain#max_allowed_rule_level_for_matching}.
	// Experimental.
	MaxAllowedRuleLevelForMatching *float64 `field:"optional" json:"maxAllowedRuleLevelForMatching" yaml:"maxAllowedRuleLevelForMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#max_allowed_rule_level_for_merging AwsDomain#max_allowed_rule_level_for_merging}.
	// Experimental.
	MaxAllowedRuleLevelForMerging *float64 `field:"optional" json:"maxAllowedRuleLevelForMerging" yaml:"maxAllowedRuleLevelForMerging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#status AwsDomain#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

