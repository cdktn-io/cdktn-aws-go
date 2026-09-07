package networkmanager


// Experimental.
type DataAwsCoreNetworkPolicyDocument_CoreNetworkConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#asn_ranges DataAwsCoreNetworkPolicyDocument#asn_ranges}.
	// Experimental.
	AsnRanges *[]*string `field:"required" json:"asnRanges" yaml:"asnRanges"`
	// edge_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#edge_locations DataAwsCoreNetworkPolicyDocument#edge_locations}
	// Experimental.
	EdgeLocations interface{} `field:"required" json:"edgeLocations" yaml:"edgeLocations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#dns_support DataAwsCoreNetworkPolicyDocument#dns_support}.
	// Experimental.
	DnsSupport interface{} `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#inside_cidr_blocks DataAwsCoreNetworkPolicyDocument#inside_cidr_blocks}.
	// Experimental.
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#security_group_referencing_support DataAwsCoreNetworkPolicyDocument#security_group_referencing_support}.
	// Experimental.
	SecurityGroupReferencingSupport interface{} `field:"optional" json:"securityGroupReferencingSupport" yaml:"securityGroupReferencingSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#vpn_ecmp_support DataAwsCoreNetworkPolicyDocument#vpn_ecmp_support}.
	// Experimental.
	VpnEcmpSupport interface{} `field:"optional" json:"vpnEcmpSupport" yaml:"vpnEcmpSupport"`
}

