package awswaf


// Experimental.
type TfWebAclRuleGroupAssociation_AwsManagedRulesAtpRuleSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#login_path TfWebAclRuleGroupAssociation#login_path}.
	// Experimental.
	LoginPath *string `field:"required" json:"loginPath" yaml:"loginPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#enable_regex_in_path TfWebAclRuleGroupAssociation#enable_regex_in_path}.
	// Experimental.
	EnableRegexInPath interface{} `field:"optional" json:"enableRegexInPath" yaml:"enableRegexInPath"`
	// request_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#request_inspection TfWebAclRuleGroupAssociation#request_inspection}
	// Experimental.
	RequestInspection interface{} `field:"optional" json:"requestInspection" yaml:"requestInspection"`
	// response_inspection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#response_inspection TfWebAclRuleGroupAssociation#response_inspection}
	// Experimental.
	ResponseInspection interface{} `field:"optional" json:"responseInspection" yaml:"responseInspection"`
}

