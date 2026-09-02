package awsdlm


// Experimental.
type TfLifecyclePolicy_ScheduleProperty struct {
	// create_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#create_rule TfLifecyclePolicy#create_rule}
	// Experimental.
	CreateRule *TfLifecyclePolicy_CreateRuleProperty `field:"required" json:"createRule" yaml:"createRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#name TfLifecyclePolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#retain_rule TfLifecyclePolicy#retain_rule}
	// Experimental.
	RetainRule *TfLifecyclePolicy_PolicyDetailsScheduleRetainRuleProperty `field:"required" json:"retainRule" yaml:"retainRule"`
	// archive_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#archive_rule TfLifecyclePolicy#archive_rule}
	// Experimental.
	ArchiveRule *TfLifecyclePolicy_ArchiveRuleProperty `field:"optional" json:"archiveRule" yaml:"archiveRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#copy_tags TfLifecyclePolicy#copy_tags}.
	// Experimental.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// cross_region_copy_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#cross_region_copy_rule TfLifecyclePolicy#cross_region_copy_rule}
	// Experimental.
	CrossRegionCopyRule interface{} `field:"optional" json:"crossRegionCopyRule" yaml:"crossRegionCopyRule"`
	// deprecate_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#deprecate_rule TfLifecyclePolicy#deprecate_rule}
	// Experimental.
	DeprecateRule *TfLifecyclePolicy_PolicyDetailsScheduleDeprecateRuleProperty `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// fast_restore_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#fast_restore_rule TfLifecyclePolicy#fast_restore_rule}
	// Experimental.
	FastRestoreRule *TfLifecyclePolicy_FastRestoreRuleProperty `field:"optional" json:"fastRestoreRule" yaml:"fastRestoreRule"`
	// share_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#share_rule TfLifecyclePolicy#share_rule}
	// Experimental.
	ShareRule *TfLifecyclePolicy_ShareRuleProperty `field:"optional" json:"shareRule" yaml:"shareRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#tags_to_add TfLifecyclePolicy#tags_to_add}.
	// Experimental.
	TagsToAdd *map[string]*string `field:"optional" json:"tagsToAdd" yaml:"tagsToAdd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#variable_tags TfLifecyclePolicy#variable_tags}.
	// Experimental.
	VariableTags *map[string]*string `field:"optional" json:"variableTags" yaml:"variableTags"`
}

