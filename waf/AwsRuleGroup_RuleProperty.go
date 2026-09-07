package waf


// Experimental.
type AwsRuleGroup_RuleProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#action AwsRuleGroup#action}
	// Experimental.
	Action *AwsRuleGroup_ActionProperty `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#name AwsRuleGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#priority AwsRuleGroup#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#visibility_config AwsRuleGroup#visibility_config}
	// Experimental.
	VisibilityConfig *AwsRuleGroup_RuleVisibilityConfigProperty `field:"required" json:"visibilityConfig" yaml:"visibilityConfig"`
	// captcha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#captcha_config AwsRuleGroup#captcha_config}
	// Experimental.
	CaptchaConfig *AwsRuleGroup_CaptchaConfigProperty `field:"optional" json:"captchaConfig" yaml:"captchaConfig"`
	// rule_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#rule_label AwsRuleGroup#rule_label}
	// Experimental.
	RuleLabel interface{} `field:"optional" json:"ruleLabel" yaml:"ruleLabel"`
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#statement AwsRuleGroup#statement}
	// Experimental.
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}

