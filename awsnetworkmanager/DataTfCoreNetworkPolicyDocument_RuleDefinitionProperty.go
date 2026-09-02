package awsnetworkmanager


// Experimental.
type DataTfCoreNetworkPolicyDocument_RuleDefinitionProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#action DataTfCoreNetworkPolicyDocument#action}
	// Experimental.
	Action *DataTfCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionProperty `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#condition_logic DataTfCoreNetworkPolicyDocument#condition_logic}.
	// Experimental.
	ConditionLogic *string `field:"optional" json:"conditionLogic" yaml:"conditionLogic"`
	// match_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#match_conditions DataTfCoreNetworkPolicyDocument#match_conditions}
	// Experimental.
	MatchConditions interface{} `field:"optional" json:"matchConditions" yaml:"matchConditions"`
}

