package awswaf


// Experimental.
type TfRuleGroup_CaptchaConfigProperty struct {
	// immunity_time_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#immunity_time_property TfRuleGroup#immunity_time_property}
	// Experimental.
	ImmunityTimeProperty *TfRuleGroup_ImmunityTimePropertyProperty `field:"optional" json:"immunityTimeProperty" yaml:"immunityTimeProperty"`
}

