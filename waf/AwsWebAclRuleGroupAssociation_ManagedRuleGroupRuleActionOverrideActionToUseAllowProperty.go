package waf


// Experimental.
type AwsWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseAllowProperty struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#custom_request_handling AwsWebAclRuleGroupAssociation#custom_request_handling}
	// Experimental.
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

