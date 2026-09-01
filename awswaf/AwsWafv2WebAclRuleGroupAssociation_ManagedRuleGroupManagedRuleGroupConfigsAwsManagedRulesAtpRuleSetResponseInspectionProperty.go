package awswaf


// Experimental.
type AwsWafv2WebAclRuleGroupAssociation_ManagedRuleGroupManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionProperty struct {
	// body_contains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#body_contains AwsWafv2WebAclRuleGroupAssociation#body_contains}
	// Experimental.
	BodyContains interface{} `field:"optional" json:"bodyContains" yaml:"bodyContains"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#header AwsWafv2WebAclRuleGroupAssociation#header}
	// Experimental.
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#json AwsWafv2WebAclRuleGroupAssociation#json}
	// Experimental.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// status_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#status_code AwsWafv2WebAclRuleGroupAssociation#status_code}
	// Experimental.
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

