package awsnetworkmanager


// Experimental.
type DataTfCoreNetworkPolicyDocument_EdgeLocationAssociationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#edge_location DataTfCoreNetworkPolicyDocument#edge_location}.
	// Experimental.
	EdgeLocation *string `field:"required" json:"edgeLocation" yaml:"edgeLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#peer_edge_location DataTfCoreNetworkPolicyDocument#peer_edge_location}.
	// Experimental.
	PeerEdgeLocation *string `field:"required" json:"peerEdgeLocation" yaml:"peerEdgeLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_names DataTfCoreNetworkPolicyDocument#routing_policy_names}.
	// Experimental.
	RoutingPolicyNames *[]*string `field:"required" json:"routingPolicyNames" yaml:"routingPolicyNames"`
}

