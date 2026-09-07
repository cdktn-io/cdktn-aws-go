package waf


// Experimental.
type AwsRuleGroup_CaptchaConfigProperty struct {
	// immunity_time_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#immunity_time_property AwsRuleGroup#immunity_time_property}
	// Experimental.
	ImmunityTimeProperty *AwsRuleGroup_ImmunityTimePropertyProperty `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

