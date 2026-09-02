package awslightsail


// Experimental.
type TfDistribution_ForwardedCookiesProperty struct {
	// The specific cookies to forward to your distribution's origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#cookies_allow_list TfDistribution#cookies_allow_list}
	// Experimental.
	CookiesAllowList *[]*string `field:"optional" json:"cookiesAllowList" yaml:"cookiesAllowList"`
	// Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#option TfDistribution#option}
	// Experimental.
	Option *string `field:"optional" json:"option" yaml:"option"`
}

