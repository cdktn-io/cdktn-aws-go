package awscloudfront


// Experimental.
type TfCachePolicy_CookiesConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#cookie_behavior TfCachePolicy#cookie_behavior}.
	// Experimental.
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_cache_policy#cookies TfCachePolicy#cookies}
	// Experimental.
	Cookies *TfCachePolicy_CookiesProperty `field:"optional" json:"cookies" yaml:"cookies"`
}

