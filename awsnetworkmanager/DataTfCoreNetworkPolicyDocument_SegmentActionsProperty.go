package awsnetworkmanager


// Experimental.
type DataTfCoreNetworkPolicyDocument_SegmentActionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#action DataTfCoreNetworkPolicyDocument#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#segment DataTfCoreNetworkPolicyDocument#segment}.
	// Experimental.
	Segment *string `field:"required" json:"segment" yaml:"segment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#description DataTfCoreNetworkPolicyDocument#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#destination_cidr_blocks DataTfCoreNetworkPolicyDocument#destination_cidr_blocks}.
	// Experimental.
	DestinationCidrBlocks *[]*string `field:"optional" json:"destinationCidrBlocks" yaml:"destinationCidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#destinations DataTfCoreNetworkPolicyDocument#destinations}.
	// Experimental.
	Destinations *[]*string `field:"optional" json:"destinations" yaml:"destinations"`
	// edge_location_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#edge_location_association DataTfCoreNetworkPolicyDocument#edge_location_association}
	// Experimental.
	EdgeLocationAssociation *DataTfCoreNetworkPolicyDocument_EdgeLocationAssociationProperty `field:"optional" json:"edgeLocationAssociation" yaml:"edgeLocationAssociation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#mode DataTfCoreNetworkPolicyDocument#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_names DataTfCoreNetworkPolicyDocument#routing_policy_names}.
	// Experimental.
	RoutingPolicyNames *[]*string `field:"optional" json:"routingPolicyNames" yaml:"routingPolicyNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#share_with DataTfCoreNetworkPolicyDocument#share_with}.
	// Experimental.
	ShareWith *[]*string `field:"optional" json:"shareWith" yaml:"shareWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#share_with_except DataTfCoreNetworkPolicyDocument#share_with_except}.
	// Experimental.
	ShareWithExcept *[]*string `field:"optional" json:"shareWithExcept" yaml:"shareWithExcept"`
	// via block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#via DataTfCoreNetworkPolicyDocument#via}
	// Experimental.
	Via *DataTfCoreNetworkPolicyDocument_ViaProperty `field:"optional" json:"via" yaml:"via"`
	// when_sent_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#when_sent_to DataTfCoreNetworkPolicyDocument#when_sent_to}
	// Experimental.
	WhenSentTo *DataTfCoreNetworkPolicyDocument_WhenSentToProperty `field:"optional" json:"whenSentTo" yaml:"whenSentTo"`
}

