package awsvpclattice


// Experimental.
type AwsVpclatticeServiceNetworkVpcAssociation_DnsOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#private_dns_preference AwsVpclatticeServiceNetworkVpcAssociation#private_dns_preference}.
	// Experimental.
	PrivateDnsPreference *string `field:"optional" json:"privateDnsPreference" yaml:"privateDnsPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_service_network_vpc_association#private_dns_specified_domains AwsVpclatticeServiceNetworkVpcAssociation#private_dns_specified_domains}.
	// Experimental.
	PrivateDnsSpecifiedDomains *[]*string `field:"optional" json:"privateDnsSpecifiedDomains" yaml:"privateDnsSpecifiedDomains"`
}

