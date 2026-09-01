package awswaf


// Experimental.
type AwsWafv2RuleGroup_CaptchaConfigProperty struct {
	// immunity_time_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#immunity_time_property AwsWafv2RuleGroup#immunity_time_property}
	// Experimental.
	ImmunityTimeProperty *AwsWafv2RuleGroup_ImmunityTimePropertyProperty `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

