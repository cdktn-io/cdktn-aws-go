package awscloudfront


// Experimental.
type AwsCloudfrontResponseHeadersPolicy_SecurityHeadersConfigProperty struct {
	// content_security_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#content_security_policy AwsCloudfrontResponseHeadersPolicy#content_security_policy}
	// Experimental.
	ContentSecurityPolicy *AwsCloudfrontResponseHeadersPolicy_ContentSecurityPolicyProperty `field:"optional" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// content_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#content_type_options AwsCloudfrontResponseHeadersPolicy#content_type_options}
	// Experimental.
	ContentTypeOptions *AwsCloudfrontResponseHeadersPolicy_ContentTypeOptionsProperty `field:"optional" json:"contentTypeOptions" yaml:"contentTypeOptions"`
	// frame_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#frame_options AwsCloudfrontResponseHeadersPolicy#frame_options}
	// Experimental.
	FrameOptions *AwsCloudfrontResponseHeadersPolicy_FrameOptionsProperty `field:"optional" json:"frameOptions" yaml:"frameOptions"`
	// referrer_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#referrer_policy AwsCloudfrontResponseHeadersPolicy#referrer_policy}
	// Experimental.
	ReferrerPolicy *AwsCloudfrontResponseHeadersPolicy_ReferrerPolicyProperty `field:"optional" json:"referrerPolicy" yaml:"referrerPolicy"`
	// strict_transport_security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#strict_transport_security AwsCloudfrontResponseHeadersPolicy#strict_transport_security}
	// Experimental.
	StrictTransportSecurity *AwsCloudfrontResponseHeadersPolicy_StrictTransportSecurityProperty `field:"optional" json:"strictTransportSecurity" yaml:"strictTransportSecurity"`
	// xss_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#xss_protection AwsCloudfrontResponseHeadersPolicy#xss_protection}
	// Experimental.
	XssProtection *AwsCloudfrontResponseHeadersPolicy_XssProtectionProperty `field:"optional" json:"xssProtection" yaml:"xssProtection"`
}

