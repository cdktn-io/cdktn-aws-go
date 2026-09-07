package networkmanager


// Experimental.
type DataAwsCoreNetworkPolicyDocument_AttachmentPoliciesActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#add_to_network_function_group DataAwsCoreNetworkPolicyDocument#add_to_network_function_group}.
	// Experimental.
	AddToNetworkFunctionGroup *string `field:"optional" json:"addToNetworkFunctionGroup" yaml:"addToNetworkFunctionGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#association_method DataAwsCoreNetworkPolicyDocument#association_method}.
	// Experimental.
	AssociationMethod *string `field:"optional" json:"associationMethod" yaml:"associationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#require_acceptance DataAwsCoreNetworkPolicyDocument#require_acceptance}.
	// Experimental.
	RequireAcceptance interface{} `field:"optional" json:"requireAcceptance" yaml:"requireAcceptance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#segment DataAwsCoreNetworkPolicyDocument#segment}.
	// Experimental.
	Segment *string `field:"optional" json:"segment" yaml:"segment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#tag_value_of_key DataAwsCoreNetworkPolicyDocument#tag_value_of_key}.
	// Experimental.
	TagValueOfKey *string `field:"optional" json:"tagValueOfKey" yaml:"tagValueOfKey"`
}

