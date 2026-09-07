package networkmanager


// Experimental.
type DataAwsCoreNetworkPolicyDocument_AttachmentPoliciesProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#action DataAwsCoreNetworkPolicyDocument#action}
	// Experimental.
	Action *DataAwsCoreNetworkPolicyDocument_AttachmentPoliciesActionProperty `field:"required" json:"action" yaml:"action"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#conditions DataAwsCoreNetworkPolicyDocument#conditions}
	// Experimental.
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#rule_number DataAwsCoreNetworkPolicyDocument#rule_number}.
	// Experimental.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#condition_logic DataAwsCoreNetworkPolicyDocument#condition_logic}.
	// Experimental.
	ConditionLogic *string `field:"optional" json:"conditionLogic" yaml:"conditionLogic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#description DataAwsCoreNetworkPolicyDocument#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

