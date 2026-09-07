package networkmanager


// Experimental.
type DataAwsCoreNetworkPolicyDocument_RuleDefinitionProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#action DataAwsCoreNetworkPolicyDocument#action}
	// Experimental.
	Action *DataAwsCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionProperty `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#condition_logic DataAwsCoreNetworkPolicyDocument#condition_logic}.
	// Experimental.
	ConditionLogic *string `field:"optional" json:"conditionLogic" yaml:"conditionLogic"`
	// match_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#match_conditions DataAwsCoreNetworkPolicyDocument#match_conditions}
	// Experimental.
	MatchConditions interface{} `field:"optional" json:"matchConditions" yaml:"matchConditions"`
}

