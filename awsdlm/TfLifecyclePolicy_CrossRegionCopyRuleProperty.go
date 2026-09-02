package awsdlm


// Experimental.
type TfLifecyclePolicy_CrossRegionCopyRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#encrypted TfLifecyclePolicy#encrypted}.
	// Experimental.
	Encrypted interface{} `field:"required" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#cmk_arn TfLifecyclePolicy#cmk_arn}.
	// Experimental.
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#copy_tags TfLifecyclePolicy#copy_tags}.
	// Experimental.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// deprecate_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#deprecate_rule TfLifecyclePolicy#deprecate_rule}
	// Experimental.
	DeprecateRule *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleProperty `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// retain_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#retain_rule TfLifecyclePolicy#retain_rule}
	// Experimental.
	RetainRule *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRuleProperty `field:"optional" json:"retainRule" yaml:"retainRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#target TfLifecyclePolicy#target}.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#target_region TfLifecyclePolicy#target_region}.
	// Experimental.
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
}

