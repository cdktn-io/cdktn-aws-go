package awsnetworkmanager


// Experimental.
type DataAwsNetworkmanagerCoreNetworkPolicyDocument_ViaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#network_function_groups DataAwsNetworkmanagerCoreNetworkPolicyDocument#network_function_groups}.
	// Experimental.
	NetworkFunctionGroups *[]*string `field:"optional" json:"networkFunctionGroups" yaml:"networkFunctionGroups"`
	// with_edge_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#with_edge_override DataAwsNetworkmanagerCoreNetworkPolicyDocument#with_edge_override}
	// Experimental.
	WithEdgeOverride interface{} `field:"optional" json:"withEdgeOverride" yaml:"withEdgeOverride"`
}

