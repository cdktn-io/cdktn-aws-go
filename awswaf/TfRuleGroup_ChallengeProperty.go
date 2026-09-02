package awswaf


// Experimental.
type TfRuleGroup_ChallengeProperty struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#custom_request_handling TfRuleGroup#custom_request_handling}
	// Experimental.
	CustomRequestHandling *TfRuleGroup_RuleActionChallengeCustomRequestHandlingProperty `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

