package awsnetworkmanager


// Experimental.
type DataTfCoreNetworkPolicyDocument_WithEdgeOverrideProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#edge_sets DataTfCoreNetworkPolicyDocument#edge_sets}.
	// Experimental.
	EdgeSets interface{} `field:"optional" json:"edgeSets" yaml:"edgeSets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#use_edge DataTfCoreNetworkPolicyDocument#use_edge}.
	// Experimental.
	UseEdge *string `field:"optional" json:"useEdge" yaml:"useEdge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#use_edge_location DataTfCoreNetworkPolicyDocument#use_edge_location}.
	// Experimental.
	UseEdgeLocation *string `field:"optional" json:"useEdgeLocation" yaml:"useEdgeLocation"`
}

