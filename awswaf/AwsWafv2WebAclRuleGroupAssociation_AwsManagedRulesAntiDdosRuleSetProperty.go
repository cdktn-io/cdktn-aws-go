package awswaf


// Experimental.
type AwsWafv2WebAclRuleGroupAssociation_AwsManagedRulesAntiDdosRuleSetProperty struct {
	// client_side_action_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#client_side_action_config AwsWafv2WebAclRuleGroupAssociation#client_side_action_config}
	// Experimental.
	ClientSideActionConfig interface{} `field:"optional" json:"clientSideActionConfig" yaml:"clientSideActionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule_group_association#sensitivity_to_block AwsWafv2WebAclRuleGroupAssociation#sensitivity_to_block}.
	// Experimental.
	SensitivityToBlock *string `field:"optional" json:"sensitivityToBlock" yaml:"sensitivityToBlock"`
}

