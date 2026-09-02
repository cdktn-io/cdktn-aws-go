package awsvpc


// Experimental.
type TfEndpoint_DnsOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#dns_record_ip_type TfEndpoint#dns_record_ip_type}.
	// Experimental.
	DnsRecordIpType *string `field:"optional" json:"dnsRecordIpType" yaml:"dnsRecordIpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#private_dns_only_for_inbound_resolver_endpoint TfEndpoint#private_dns_only_for_inbound_resolver_endpoint}.
	// Experimental.
	PrivateDnsOnlyForInboundResolverEndpoint interface{} `field:"optional" json:"privateDnsOnlyForInboundResolverEndpoint" yaml:"privateDnsOnlyForInboundResolverEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#private_dns_preference TfEndpoint#private_dns_preference}.
	// Experimental.
	PrivateDnsPreference *string `field:"optional" json:"privateDnsPreference" yaml:"privateDnsPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#private_dns_specified_domains TfEndpoint#private_dns_specified_domains}.
	// Experimental.
	PrivateDnsSpecifiedDomains *[]*string `field:"optional" json:"privateDnsSpecifiedDomains" yaml:"privateDnsSpecifiedDomains"`
}

