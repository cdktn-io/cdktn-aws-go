package awsnetworkmanager


// Experimental.
type DataTfCoreNetworkPolicyDocument_RoutingPolicyRulesProperty struct {
	// rule_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#rule_definition DataTfCoreNetworkPolicyDocument#rule_definition}
	// Experimental.
	RuleDefinition *DataTfCoreNetworkPolicyDocument_RuleDefinitionProperty `field:"required" json:"ruleDefinition" yaml:"ruleDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#rule_number DataTfCoreNetworkPolicyDocument#rule_number}.
	// Experimental.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
}

