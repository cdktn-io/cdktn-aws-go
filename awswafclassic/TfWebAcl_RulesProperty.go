package awswafclassic


// Experimental.
type TfWebAcl_RulesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_web_acl#priority TfWebAcl#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_web_acl#rule_id TfWebAcl#rule_id}.
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_web_acl#action TfWebAcl#action}
	// Experimental.
	Action *TfWebAcl_ActionProperty `field:"optional" json:"action" yaml:"action"`
	// override_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_web_acl#override_action TfWebAcl#override_action}
	// Experimental.
	OverrideAction *TfWebAcl_OverrideActionProperty `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_web_acl#type TfWebAcl#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

