package waf


// Experimental.
type AwsRuleGroup_CaptchaProperty struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#custom_request_handling AwsRuleGroup#custom_request_handling}
	// Experimental.
	CustomRequestHandling *AwsRuleGroup_RuleActionCaptchaCustomRequestHandlingProperty `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

