package awsnetworkmanager


// Experimental.
type DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#action DataTfCoreNetworkPolicyDocument#action}
	// Experimental.
	Action *DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesActionProperty `field:"required" json:"action" yaml:"action"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#conditions DataTfCoreNetworkPolicyDocument#conditions}
	// Experimental.
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#rule_number DataTfCoreNetworkPolicyDocument#rule_number}.
	// Experimental.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#description DataTfCoreNetworkPolicyDocument#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#edge_locations DataTfCoreNetworkPolicyDocument#edge_locations}.
	// Experimental.
	EdgeLocations *[]*string `field:"optional" json:"edgeLocations" yaml:"edgeLocations"`
}

