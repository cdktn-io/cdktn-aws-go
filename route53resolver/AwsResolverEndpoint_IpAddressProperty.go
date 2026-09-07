package route53resolver


// Experimental.
type AwsResolverEndpoint_IpAddressProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#subnet_id AwsResolverEndpoint#subnet_id}.
	// Experimental.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#ip AwsResolverEndpoint#ip}.
	// Experimental.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#ipv6 AwsResolverEndpoint#ipv6}.
	// Experimental.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
}

