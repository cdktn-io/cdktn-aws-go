package awscloudfront


// Experimental.
type TfOriginRequestPolicy_CookiesConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_origin_request_policy#cookie_behavior TfOriginRequestPolicy#cookie_behavior}.
	// Experimental.
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_origin_request_policy#cookies TfOriginRequestPolicy#cookies}
	// Experimental.
	Cookies *TfOriginRequestPolicy_CookiesProperty `field:"optional" json:"cookies" yaml:"cookies"`
}

