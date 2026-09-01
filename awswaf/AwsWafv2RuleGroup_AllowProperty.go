package awswaf


// Experimental.
type AwsWafv2RuleGroup_AllowProperty struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#custom_request_handling AwsWafv2RuleGroup#custom_request_handling}
	// Experimental.
	CustomRequestHandling *AwsWafv2RuleGroup_RuleActionAllowCustomRequestHandlingProperty `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

