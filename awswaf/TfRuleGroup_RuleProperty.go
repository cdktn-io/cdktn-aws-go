package awswaf


// Experimental.
type TfRuleGroup_RuleProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#action TfRuleGroup#action}
	// Experimental.
	Action *TfRuleGroup_ActionProperty `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#name TfRuleGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#priority TfRuleGroup#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#visibility_config TfRuleGroup#visibility_config}
	// Experimental.
	VisibilityConfig *TfRuleGroup_RuleVisibilityConfigProperty `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// captcha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#captcha_config TfRuleGroup#captcha_config}
	// Experimental.
	CaptchaConfig *TfRuleGroup_CaptchaConfigProperty `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// rule_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#rule_label TfRuleGroup#rule_label}
	// Experimental.
	RuleLabel interface{} `field:"optional" json:"ruleLabel" yaml:"ruleLabel"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#statement TfRuleGroup#statement}
	// Experimental.
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}

