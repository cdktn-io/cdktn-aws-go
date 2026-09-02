package awsconnectcustomerprofiles


// Experimental.
type TfDomain_RuleBasedMatchingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#enabled TfDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// attribute_types_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#attribute_types_selector TfDomain#attribute_types_selector}
	// Experimental.
	AttributeTypesSelector *TfDomain_AttributeTypesSelectorProperty `field:"optional" json:"attributeTypesSelector" yaml:"attributeTypesSelector"`
	// conflict_resolution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#conflict_resolution TfDomain#conflict_resolution}
	// Experimental.
	ConflictResolution *TfDomain_RuleBasedMatchingConflictResolutionProperty `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// exporting_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#exporting_config TfDomain#exporting_config}
	// Experimental.
	ExportingConfig *TfDomain_RuleBasedMatchingExportingConfigProperty `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// matching_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#matching_rules TfDomain#matching_rules}
	// Experimental.
	MatchingRules interface{} `field:"optional" json:"matchingRules" yaml:"matchingRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#max_allowed_rule_level_for_matching TfDomain#max_allowed_rule_level_for_matching}.
	// Experimental.
	MaxAllowedRuleLevelForMatching *float64 `field:"optional" json:"maxAllowedRuleLevelForMatching" yaml:"maxAllowedRuleLevelForMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#max_allowed_rule_level_for_merging TfDomain#max_allowed_rule_level_for_merging}.
	// Experimental.
	MaxAllowedRuleLevelForMerging *float64 `field:"optional" json:"maxAllowedRuleLevelForMerging" yaml:"maxAllowedRuleLevelForMerging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#status TfDomain#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

