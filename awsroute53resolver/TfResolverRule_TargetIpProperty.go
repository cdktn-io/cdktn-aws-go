package awsroute53resolver


// Experimental.
type TfResolverRule_TargetIpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_rule#ip TfResolverRule#ip}.
	// Experimental.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_rule#ipv6 TfResolverRule#ipv6}.
	// Experimental.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_rule#port TfResolverRule#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_rule#protocol TfResolverRule#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

