package awslightsail


// Experimental.
type AwsLightsailDistribution_ForwardedHeadersProperty struct {
	// The specific headers to forward to your distribution's origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#headers_allow_list AwsLightsailDistribution#headers_allow_list}
	// Experimental.
	HeadersAllowList *[]*string `field:"optional" json:"headersAllowList" yaml:"headersAllowList"`
	// The headers that you want your distribution to forward to your origin and base caching on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#option AwsLightsailDistribution#option}
	// Experimental.
	Option *string `field:"optional" json:"option" yaml:"option"`
}

