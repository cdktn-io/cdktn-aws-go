package cloudfront


// Experimental.
type AwsResponseHeadersPolicy_SecurityHeadersConfigProperty struct {
	// content_security_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#content_security_policy AwsResponseHeadersPolicy#content_security_policy}
	// Experimental.
	ContentSecurityPolicy *AwsResponseHeadersPolicy_ContentSecurityPolicyProperty `field:"optional" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// content_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#content_type_options AwsResponseHeadersPolicy#content_type_options}
	// Experimental.
	ContentTypeOptions *AwsResponseHeadersPolicy_ContentTypeOptionsProperty `field:"optional" json:"contentTypeOptions" yaml:"contentTypeOptions"`
	// frame_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#frame_options AwsResponseHeadersPolicy#frame_options}
	// Experimental.
	FrameOptions *AwsResponseHeadersPolicy_FrameOptionsProperty `field:"optional" json:"frameOptions" yaml:"frameOptions"`
	// referrer_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#referrer_policy AwsResponseHeadersPolicy#referrer_policy}
	// Experimental.
	ReferrerPolicy *AwsResponseHeadersPolicy_ReferrerPolicyProperty `field:"optional" json:"referrerPolicy" yaml:"referrerPolicy"`
	// strict_transport_security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#strict_transport_security AwsResponseHeadersPolicy#strict_transport_security}
	// Experimental.
	StrictTransportSecurity *AwsResponseHeadersPolicy_StrictTransportSecurityProperty `field:"optional" json:"strictTransportSecurity" yaml:"strictTransportSecurity"`
	// xss_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#xss_protection AwsResponseHeadersPolicy#xss_protection}
	// Experimental.
	XssProtection *AwsResponseHeadersPolicy_XssProtectionProperty `field:"optional" json:"xssProtection" yaml:"xssProtection"`
}

