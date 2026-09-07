package networkmanager


// Experimental.
type DataAwsCoreNetworkPolicyDocument_EdgeLocationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#location DataAwsCoreNetworkPolicyDocument#location}.
	// Experimental.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#asn DataAwsCoreNetworkPolicyDocument#asn}.
	// Experimental.
	Asn *string `field:"optional" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#inside_cidr_blocks DataAwsCoreNetworkPolicyDocument#inside_cidr_blocks}.
	// Experimental.
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
}

