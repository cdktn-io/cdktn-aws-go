package awsnetworkmanager


// Experimental.
type DataAwsNetworkmanagerCoreNetworkPolicyDocument_AttachmentPoliciesActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#add_to_network_function_group DataAwsNetworkmanagerCoreNetworkPolicyDocument#add_to_network_function_group}.
	// Experimental.
	AddToNetworkFunctionGroup *string `field:"optional" json:"addToNetworkFunctionGroup" yaml:"addToNetworkFunctionGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#association_method DataAwsNetworkmanagerCoreNetworkPolicyDocument#association_method}.
	// Experimental.
	AssociationMethod *string `field:"optional" json:"associationMethod" yaml:"associationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#require_acceptance DataAwsNetworkmanagerCoreNetworkPolicyDocument#require_acceptance}.
	// Experimental.
	RequireAcceptance interface{} `field:"optional" json:"requireAcceptance" yaml:"requireAcceptance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#segment DataAwsNetworkmanagerCoreNetworkPolicyDocument#segment}.
	// Experimental.
	Segment *string `field:"optional" json:"segment" yaml:"segment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#tag_value_of_key DataAwsNetworkmanagerCoreNetworkPolicyDocument#tag_value_of_key}.
	// Experimental.
	TagValueOfKey *string `field:"optional" json:"tagValueOfKey" yaml:"tagValueOfKey"`
}

