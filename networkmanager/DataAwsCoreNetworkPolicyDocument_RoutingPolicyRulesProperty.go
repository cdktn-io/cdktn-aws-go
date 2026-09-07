package networkmanager


// Experimental.
type DataAwsCoreNetworkPolicyDocument_RoutingPolicyRulesProperty struct {
	// rule_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#rule_definition DataAwsCoreNetworkPolicyDocument#rule_definition}
	// Experimental.
	RuleDefinition *DataAwsCoreNetworkPolicyDocument_RuleDefinitionProperty `field:"required" json:"ruleDefinition" yaml:"ruleDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#rule_number DataAwsCoreNetworkPolicyDocument#rule_number}.
	// Experimental.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
}

