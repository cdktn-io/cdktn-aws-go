package awswaf


// Experimental.
type TfWebAclRule_ManagedRuleGroupStatementProperty struct {
	// Name of the managed rule group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#name TfWebAclRule#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the managed rule group vendor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#vendor_name TfWebAclRule#vendor_name}
	// Experimental.
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// managed_rule_group_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#managed_rule_group_configs TfWebAclRule#managed_rule_group_configs}
	// Experimental.
	ManagedRuleGroupConfigs interface{} `field:"optional" json:"managedRuleGroupConfigs" yaml:"managedRuleGroupConfigs"`
	// rule_action_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#rule_action_override TfWebAclRule#rule_action_override}
	// Experimental.
	RuleActionOverride interface{} `field:"optional" json:"ruleActionOverride" yaml:"ruleActionOverride"`
	// scope_down_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#scope_down_statement TfWebAclRule#scope_down_statement}
	// Experimental.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
	// Version of the managed rule group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#version TfWebAclRule#version}
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

