package awsdlm


// Experimental.
type AwsDlmLifecyclePolicy_ScheduleProperty struct {
	// create_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#create_rule AwsDlmLifecyclePolicy#create_rule}
	// Experimental.
	CreateRule *AwsDlmLifecyclePolicy_CreateRuleProperty `field:"required" json:"createRule" yaml:"createRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#name AwsDlmLifecyclePolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#retain_rule AwsDlmLifecyclePolicy#retain_rule}
	// Experimental.
	RetainRule *AwsDlmLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty `field:"required" json:"retainRule" yaml:"retainRule"`
	// archive_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#archive_rule AwsDlmLifecyclePolicy#archive_rule}
	// Experimental.
	ArchiveRule *AwsDlmLifecyclePolicy_ArchiveRuleProperty `field:"optional" json:"archiveRule" yaml:"archiveRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#copy_tags AwsDlmLifecyclePolicy#copy_tags}.
	// Experimental.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// cross_region_copy_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#cross_region_copy_rule AwsDlmLifecyclePolicy#cross_region_copy_rule}
	// Experimental.
	CrossRegionCopyRule interface{} `field:"optional" json:"crossRegionCopyRule" yaml:"crossRegionCopyRule"`
	// deprecate_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#deprecate_rule AwsDlmLifecyclePolicy#deprecate_rule}
	// Experimental.
	DeprecateRule *AwsDlmLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// fast_restore_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#fast_restore_rule AwsDlmLifecyclePolicy#fast_restore_rule}
	// Experimental.
	FastRestoreRule *AwsDlmLifecyclePolicy_FastRestoreRuleProperty `field:"optional" json:"fastRestoreRule" yaml:"fastRestoreRule"`
	// share_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#share_rule AwsDlmLifecyclePolicy#share_rule}
	// Experimental.
	ShareRule *AwsDlmLifecyclePolicy_ShareRuleProperty `field:"optional" json:"shareRule" yaml:"shareRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#tags_to_add AwsDlmLifecyclePolicy#tags_to_add}.
	// Experimental.
	TagsToAdd *map[string]*string `field:"optional" json:"tagsToAdd" yaml:"tagsToAdd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#variable_tags AwsDlmLifecyclePolicy#variable_tags}.
	// Experimental.
	VariableTags *map[string]*string `field:"optional" json:"variableTags" yaml:"variableTags"`
}

