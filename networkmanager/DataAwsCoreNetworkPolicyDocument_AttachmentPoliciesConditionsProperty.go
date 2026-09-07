package networkmanager


// Experimental.
type DataAwsCoreNetworkPolicyDocument_AttachmentPoliciesConditionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#type DataAwsCoreNetworkPolicyDocument#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#key DataAwsCoreNetworkPolicyDocument#key}.
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#operator DataAwsCoreNetworkPolicyDocument#operator}.
	// Experimental.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#value DataAwsCoreNetworkPolicyDocument#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

