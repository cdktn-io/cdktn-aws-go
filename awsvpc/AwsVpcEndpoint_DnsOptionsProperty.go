package awsvpc


// Experimental.
type AwsVpcEndpoint_DnsOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#dns_record_ip_type AwsVpcEndpoint#dns_record_ip_type}.
	// Experimental.
	DnsRecordIpType *string `field:"optional" json:"dnsRecordIpType" yaml:"dnsRecordIpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#private_dns_only_for_inbound_resolver_endpoint AwsVpcEndpoint#private_dns_only_for_inbound_resolver_endpoint}.
	// Experimental.
	PrivateDnsOnlyForInboundResolverEndpoint interface{} `field:"optional" json:"privateDnsOnlyForInboundResolverEndpoint" yaml:"privateDnsOnlyForInboundResolverEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#private_dns_preference AwsVpcEndpoint#private_dns_preference}.
	// Experimental.
	PrivateDnsPreference *string `field:"optional" json:"privateDnsPreference" yaml:"privateDnsPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#private_dns_specified_domains AwsVpcEndpoint#private_dns_specified_domains}.
	// Experimental.
	PrivateDnsSpecifiedDomains *[]*string `field:"optional" json:"privateDnsSpecifiedDomains" yaml:"privateDnsSpecifiedDomains"`
}

